// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package ipc provides an ipc client model and access to deployed smart contract calls.
package ipc

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zeebo/errs"

	"github.com/akave-ai/akavesdk/private/ipc/contracts"
)

// Config represents configuration for the storage contract client.
type Config struct {
	DialURI                        string        `usage:"addr of ipc endpoint"`
	PrivateKey                     string        `usage:"hex private key used to sign transactions"`
	StorageContractAddress         string        `usage:"hex storage contract address"`
	AccessContractAddress          string        `usage:"hex access manager contract address"`
	PolicyFactoryContractAddress   string        `usage:"hex policy factory contract address"`
	UseCustomValues                bool          `usage:"use custom blockchain values (gas price, gas limit, nonce)"`
	GasPriceCheckInterval          time.Duration `usage:"gas price check loop renewal interval"`
	NumberOfConcurrentTransactions int64         `usage:"number of concurrent fill chunk block transactions"`
}

// StorageData represents the struct for signing.
type StorageData struct {
	ChunkCID   []byte
	BlockCID   [32]byte
	ChunkIndex *big.Int
	BlockIndex uint8
	NodeID     [32]byte
	Nonce      *big.Int
	Deadline   *big.Int
	BucketID   [32]byte
}

// DefaultConfig returns default configuration for the ipc.
func DefaultConfig() Config {
	return Config{
		DialURI:                        "",
		PrivateKey:                     "",
		StorageContractAddress:         "",
		AccessContractAddress:          "",
		PolicyFactoryContractAddress:   "",
		UseCustomValues:                false,
		GasPriceCheckInterval:          5 * time.Second,
		NumberOfConcurrentTransactions: 10,
	}
}

// Client represents storage client.
type Client struct {
	Storage          *contracts.Storage
	AccessManager    *contracts.AccessManager
	PolicyFactory    *contracts.PolicyFactory
	ListPolicyAbi    *abi.ABI
	PolicyFactoryAbi *abi.ABI
	Auth             *bind.TransactOpts
	Eth              *ethclient.Client
	chainID          *big.Int
	ticker           *time.Ticker
}

// Dial creates eth client, new smart-contract instance, auth.
func Dial(ctx context.Context, config Config) (*Client, error) {
	client, err := ethclient.Dial(config.DialURI)
	if err != nil {
		return &Client{}, err
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return &Client{}, err
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return &Client{}, err
	}

	storage, err := contracts.NewStorage(common.HexToAddress(config.StorageContractAddress), client)
	if err != nil {
		return &Client{}, err
	}

	accessManager, err := contracts.NewAccessManager(common.HexToAddress(config.AccessContractAddress), client)
	if err != nil {
		return &Client{}, err
	}

	lpAbi, err := contracts.ListPolicyMetaData.GetAbi()
	if err != nil {
		return &Client{}, err
	}

	pfAbi, err := contracts.PolicyFactoryMetaData.GetAbi()
	if err != nil {
		return &Client{}, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return &Client{}, err
	}

	ipcClient := &Client{
		Storage:          storage,
		AccessManager:    accessManager,
		ListPolicyAbi:    lpAbi,
		PolicyFactoryAbi: pfAbi,
		Auth:             auth,
		chainID:          chainID,
		Eth:              client,
		ticker:           time.NewTicker(200 * time.Millisecond),
	}

	if config.PolicyFactoryContractAddress != "" {
		ipcClient.PolicyFactory, err = contracts.NewPolicyFactory(common.HexToAddress(config.PolicyFactoryContractAddress), client)
		if err != nil {
			return &Client{}, err
		}
	}

	return ipcClient, nil
}

// DeployStorage deploys storage smart contract, returns it's client.
func DeployStorage(ctx context.Context, config Config) (*Client, string, string, error) {
	ethClient, err := ethclient.Dial(config.DialURI)
	if err != nil {
		return &Client{}, "", "", err
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return &Client{}, "", "", err
	}

	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		return &Client{}, "", "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return &Client{}, "", "", err
	}

	client := &Client{
		Auth:   auth,
		Eth:    ethClient,
		ticker: time.NewTicker(200 * time.Millisecond),
	}

	akaveTokenAddr, tx, token, err := contracts.DeployAkaveToken(auth, ethClient)
	if err != nil {
		return &Client{}, "", "", err
	}

	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, "", "", err
	}

	storageAddress, tx, storage, err := contracts.DeployStorage(auth, ethClient, akaveTokenAddr)
	if err != nil {
		return &Client{}, "", "", err
	}

	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, "", "", err
	}
	client.Storage = storage
	client.chainID = chainID

	minterRole, err := token.MINTERROLE(&bind.CallOpts{})
	if err != nil {
		return &Client{}, "", "", err
	}

	tx, err = token.GrantRole(auth, minterRole, storageAddress)
	if err != nil {
		return &Client{}, "", "", err
	}

	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, "", "", err
	}

	accessAddress, tx, accessManager, err := contracts.DeployAccessManager(client.Auth, client.Eth, storageAddress)
	if err != nil {
		return &Client{}, "", "", err
	}
	client.AccessManager = accessManager

	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, "", "", err
	}

	tx, err = storage.SetAccessManager(client.Auth, accessAddress)
	if err != nil {
		return &Client{}, "", "", err
	}

	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, "", "", err
	}

	baseListPolicyAddress, tx, _, err := contracts.DeployListPolicy(client.Auth, client.Eth)
	if err != nil {
		return &Client{}, "", "", err
	}

	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, "", "", err
	}

	_, tx, client.PolicyFactory, err = contracts.DeployPolicyFactory(client.Auth, client.Eth, baseListPolicyAddress)
	if err != nil {
		return &Client{}, "", "", err
	}

	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, "", "", err
	}

	client.ListPolicyAbi, err = contracts.ListPolicyMetaData.GetAbi()
	if err != nil {
		return &Client{}, "", "", err
	}

	client.PolicyFactoryAbi, err = contracts.PolicyFactoryMetaData.GetAbi()
	if err != nil {
		return &Client{}, "", "", err
	}

	return client, storageAddress.String(), accessAddress.String(), nil
}

// ChainID returns chain id.
func (client *Client) ChainID() *big.Int {
	return client.chainID
}

// DeployListPolicy deploys new list policy for provided user address.
func (client *Client) DeployListPolicy(ctx context.Context, user common.Address) (*contracts.ListPolicy, error) {
	abiByte, err := client.ListPolicyAbi.Pack("initialize", user)
	if err != nil {
		return nil, err
	}

	tx, err := client.PolicyFactory.DeployPolicy(client.Auth, abiByte)
	if err != nil {
		return nil, err
	}

	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return nil, err
	}

	r, err := client.Eth.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return nil, err
	}

	eventHash := client.PolicyFactoryAbi.Events["PolicyDeployed"].ID
	var policyInstance common.Address
	for _, log := range r.Logs {
		if log.Topics[0] == eventHash {
			policyInstance = common.HexToAddress(log.Topics[2].Hex())
			break
		}
	}

	listPolicy, err := contracts.NewListPolicy(policyInstance, client.Eth)
	if err != nil {
		return nil, err
	}

	return listPolicy, nil
}

// BatchReceiptRequest represents a single receipt request in a batch.
type BatchReceiptRequest struct {
	Hash common.Hash
	Key  string
}

// BatchReceiptResponse represents the response for a single receipt request.
type BatchReceiptResponse struct {
	Receipt *types.Receipt
	Error   error
	Key     string
}

// BatchReceiptResult contains all responses from a batch request.
type BatchReceiptResult struct {
	Responses []BatchReceiptResponse
}

// GetTransactionReceiptsBatch fetches multiple transaction receipts in a single batch call.
func (client *Client) GetTransactionReceiptsBatch(ctx context.Context, requests []BatchReceiptRequest, timeout time.Duration) (*BatchReceiptResult, error) {
	rpcClient := client.Eth.Client()

	batchReqs := make([]rpc.BatchElem, len(requests))
	for i, req := range requests {
		batchReqs[i] = rpc.BatchElem{
			Method: "eth_getTransactionReceipt",
			Args:   []interface{}{req.Hash},
			Result: new(*types.Receipt),
		}
	}

	batchCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := rpcClient.BatchCallContext(batchCtx, batchReqs); err != nil {
		return nil, err
	}

	responses := make([]BatchReceiptResponse, len(requests))
	for i, req := range requests {
		batchElem := batchReqs[i]
		response := BatchReceiptResponse{
			Key: req.Key,
		}

		if batchElem.Error != nil {
			response.Error = batchElem.Error
		} else {
			response.Receipt = *batchElem.Result.(**types.Receipt)

			if response.Receipt == nil {
				response.Error = ethereum.NotFound
			}
		}

		responses[i] = response
	}

	return &BatchReceiptResult{Responses: responses}, nil
}

// WaitForTx block execution until transaction receipt is received or context is cancelled.
func (client *Client) WaitForTx(ctx context.Context, hash common.Hash) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-client.ticker.C:
			receipt, err := client.Eth.TransactionReceipt(ctx, hash)
			if err == nil {
				if receipt.Status == 1 {
					return nil
				}

				return errs.New("transaction failed")
			}
			if !errors.Is(err, ethereum.NotFound) {
				return err
			}
		}
	}
}
