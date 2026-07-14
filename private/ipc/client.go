// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package ipc provides an ipc client model and access to deployed smart contract calls.
package ipc

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/akave-ai/akavesdk/private/ipc/contracts"
)

// Config represents configuration for the storage contract client.
type Config struct {
	DialURI                string `usage:"addr of ipc endpoint"`
	PrivateKey             string `usage:"hex private key used to sign transactions"`
	StorageContractAddress string `usage:"hex storage contract address"`
	AccessContractAddress  string `usage:"hex access manager contract address"`
}

// StorageData represents the struct for signing.
type StorageData struct {
	ChunkCID   []byte
	BlockCID   [32]byte
	ChunkIndex *big.Int
	BlockIndex *big.Int
	NodeID     [32]byte
	Nonce      *big.Int
	Deadline   *big.Int
	BucketID   [32]byte
}

// DefaultConfig returns default configuration for the ipc.
func DefaultConfig() Config {
	return Config{
		DialURI:                "",
		PrivateKey:             "",
		StorageContractAddress: "",
		AccessContractAddress:  "",
	}
}

// Client represents storage client.
type Client struct {
	Storage       *contracts.Storage
	AccessManager *contracts.AccessManager
	Auth          *bind.TransactOpts
	Eth           *ethclient.Client
	Addresses     ContractsAddresses
	chainID       *big.Int
}

// ContractsAddresses contains addresses of deployed contracts.
type ContractsAddresses struct {
	Token         string
	StorageImpl   string
	Storage       string
	AccessManager string
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

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return &Client{}, err
	}

	ipcClient := &Client{
		Storage:       storage,
		AccessManager: accessManager,
		Auth:          auth,
		chainID:       chainID,
		Eth:           client,
		Addresses: ContractsAddresses{
			Storage:       config.StorageContractAddress,
			AccessManager: config.AccessContractAddress,
		},
	}

	return ipcClient, nil
}

// DeployContracts deploys smart contracts, returns client.
func DeployContracts(ctx context.Context, dialURI, privateKeyHex string) (*Client, error) {
	ethClient, err := ethclient.Dial(dialURI)
	if err != nil {
		return &Client{}, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return &Client{}, err
	}

	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		return &Client{}, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return &Client{}, err
	}

	client := &Client{
		Auth:    auth,
		Eth:     ethClient,
		chainID: chainID,
	}

	akaveTokenAddr, tx, token, err := contracts.DeployAkaveToken(auth, ethClient)
	if err != nil {
		return &Client{}, err
	}

	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, err
	}

	client.Addresses.Token = akaveTokenAddr.String()

	storageImplAddress, tx, _, err := contracts.DeployStorage(auth, ethClient)
	if err != nil {
		return &Client{}, err
	}
	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, err
	}

	client.Addresses.StorageImpl = storageImplAddress.String()

	storageABI, err := contracts.StorageMetaData.GetAbi()
	if err != nil {
		return &Client{}, err
	}

	initData, err := storageABI.Pack("initialize", akaveTokenAddr)
	if err != nil {
		return &Client{}, err
	}

	storageProxyAddress, tx, _, err := contracts.DeployERC1967Proxy(
		auth,
		ethClient,
		storageImplAddress,
		initData,
	)
	if err != nil {
		return &Client{}, err
	}
	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, err
	}

	storage, err := contracts.NewStorage(storageProxyAddress, ethClient)
	if err != nil {
		return &Client{}, err
	}

	client.Storage = storage
	client.Addresses.Storage = storageProxyAddress.String()

	minterRole, err := token.MINTERROLE(&bind.CallOpts{Context: ctx})
	if err != nil {
		return &Client{}, err
	}

	tx, err = token.GrantRole(auth, minterRole, storageProxyAddress)
	if err != nil {
		return &Client{}, err
	}

	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, err
	}

	accessAddress, tx, accessManager, err := contracts.DeployAccessManager(client.Auth, client.Eth, storageProxyAddress)
	if err != nil {
		return &Client{}, err
	}
	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, err
	}

	client.AccessManager = accessManager
	client.Addresses.AccessManager = accessAddress.String()

	tx, err = storage.SetAccessManager(client.Auth, accessAddress)
	if err != nil {
		return &Client{}, err
	}
	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, err
	}

	return client, nil
}

// UpgradeStorage deploys a new Storage implementation and upgrades the existing Storage proxy
// (ERC1967/UUPS) to point to it.
//
// The proxy address is provided as proxyAddr parameter.
// Existing state remains in the proxy storage and will be visible to the new implementation.
// It is assumed that generated code used in this function belongs to a new version of the Storage contract.
func UpgradeStorage(ctx context.Context, dialURI, privateKey, proxyAddr, callDataHex string) (common.Address, error) {
	ethClient, err := ethclient.Dial(dialURI)
	if err != nil {
		return common.Address{}, err
	}
	defer ethClient.Close()

	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return common.Address{}, err
	}

	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		return common.Address{}, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
	if err != nil {
		return common.Address{}, err
	}
	auth.Context = ctx

	ipcClient := &Client{
		Eth:     ethClient,
		chainID: chainID,
	}

	storageProxy, err := contracts.NewStorage(common.HexToAddress(proxyAddr), ethClient)
	if err != nil {
		return common.Address{}, err
	}

	newImplementationAddr, deployTx, _, err := contracts.DeployStorage(auth, ethClient)
	if err != nil {
		return common.Address{}, err
	}

	if err := ipcClient.WaitForTx(ctx, deployTx.Hash()); err != nil {
		return common.Address{}, err
	}

	var callData []byte
	if callDataHex != "" {
		callData, err = hex.DecodeString(callDataHex)
		if err != nil {
			return common.Address{}, err
		}
	}

	upgradeTx, err := storageProxy.UpgradeToAndCall(auth, newImplementationAddr, callData)
	if err != nil {
		return common.Address{}, err
	}

	if err := ipcClient.WaitForTx(ctx, upgradeTx.Hash()); err != nil {
		return common.Address{}, err
	}

	return newImplementationAddr, nil
}

// ChainID returns chain id.
func (client *Client) ChainID() *big.Int {
	return client.chainID
}

// TestDeployListPolicy deploys new list policy for provided user address.
func (client *Client) TestDeployListPolicy(ctx context.Context, user common.Address) (*contracts.ListPolicy, error) {
	_, tx, listPolicy, err := contracts.DeployListPolicy(client.Auth, client.Eth)
	if err != nil {
		return nil, err
	}

	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return nil, err
	}

	tx, err = listPolicy.Initialize(client.Auth, user)
	if err != nil {
		return nil, err
	}

	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return nil, err
	}

	return listPolicy, nil
}

// WaitForTx block execution until transaction receipt is received or context is cancelled.
func (client *Client) WaitForTx(ctx context.Context, hash common.Hash) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	receipt, err := client.Eth.TransactionReceipt(ctx, hash)
	if err == nil {
		if receipt.Status == 1 {
			return nil
		}

		return errors.New("transaction failed")
	}
	if !errors.Is(err, ethereum.NotFound) {
		return err
	}

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			receipt, err := client.Eth.TransactionReceipt(ctx, hash)
			if err == nil {
				if receipt.Status == 1 {
					return nil
				}

				return errors.New("transaction failed")
			}
			if !errors.Is(err, ethereum.NotFound) {
				return err
			}
		}
	}
}
