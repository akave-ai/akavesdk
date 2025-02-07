// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package ipc provides an ipc client model and access to deployed smart contract calls.
package ipc

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zeebo/errs"

	"github.com/akave-ai/akavesdk/private/ipc/contracts"
)

// Config represents configuration for the storage contract client.
type Config struct {
	DialURI                string `usage:"addr of ipc endpoint"`
	PrivateKey             string `usage:"hex private key used to sign transactions"`
	StorageContractAddress string `usage:"hex storage contract address"`
	AccessContractAddress  string `usage:"hex access manager contract address"`
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
	client        *ethclient.Client
	ticker        *time.Ticker
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

	return &Client{
		Storage:       storage,
		AccessManager: accessManager,
		Auth:          auth,
		client:        client,
		ticker:        time.NewTicker(200 * time.Millisecond),
	}, nil
}

// DeployStorage deploys storage smart contract, returns it's client.
func DeployStorage(ctx context.Context, config Config) (*Client, string, error) {
	ethClient, err := ethclient.Dial(config.DialURI)
	if err != nil {
		return &Client{}, "", err
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return &Client{}, "", err
	}

	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		return &Client{}, "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return &Client{}, "", err
	}

	address, tx, storage, err := contracts.DeployStorage(auth, ethClient)
	if err != nil {
		return &Client{}, "", err
	}

	client := &Client{
		Storage: storage,
		Auth:    auth,
		client:  ethClient,
		ticker:  time.NewTicker(200 * time.Millisecond),
	}

	if err := client.WaitForTx(ctx, tx.Hash()); err != nil {
		return &Client{}, "", err
	}

	_, tx, client.AccessManager, err = contracts.DeployAccessManager(auth, ethClient, address)
	if err != nil {
		return &Client{}, "", err
	}

	return client, address.String(), client.WaitForTx(ctx, tx.Hash())
}

// WaitForTx block execution until transaction receipt is received or context is cancelled.
func (client *Client) WaitForTx(ctx context.Context, hash common.Hash) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return errors.New("context canceled")
		case <-client.ticker.C:
			receipt, err := client.client.TransactionReceipt(ctx, hash)
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
