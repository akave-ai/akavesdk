// Copyright (C) 2025 Akave
// See LICENSE for copying information.

// Package ipctest provides an ipc testing utils.
package ipctest

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"flag"
	"math/big"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zeebo/errs"
)

var (
	// DialURI is flag to set ipc dial uri.
	DialURI = flag.String("ipc-rpc-uri", os.Getenv("DIAL_URI"), "flag to set ipc dial uri for client tests")
	// WsURI is flag to set ipc ws uri.
	WsURI = flag.String("ipc-ws-uri", os.Getenv("WS_URI"), "flag to set ipc ws uri for client tests")
	// PrivateKey is flag to set deployers hex private key.
	PrivateKey = flag.String("private-key", os.Getenv("PRIVATE_KEY"), "flag to set deployers hex private key for client tests")
)

// PickDialURI picks IPC provider URI.
func PickDialURI(t testing.TB) string {
	if *DialURI == "" || strings.EqualFold(*DialURI, "omit") {
		t.Skip("dial uri flag missing, example: -DIAL_URI=<dial uri>")
	}
	return *DialURI
}

// PickWsURI picks IPC ws URI.
func PickWsURI(t testing.TB) string {
	if *WsURI == "" || strings.EqualFold(*WsURI, "omit") {
		t.Skip("ws uri flag missing, example: -WS_URI=<ws uri>")
	}
	return *WsURI
}

// PickPrivateKey picks hex private key of deployer.
func PickPrivateKey(t testing.TB) string {
	if *PrivateKey == "" || strings.EqualFold(*PrivateKey, "omit") {
		t.Skip("private key flag missing, example: -PRIVATE_KEY=<deployers hex private key>")
	}
	return *PrivateKey
}

var mu sync.Mutex

// NewFundedAccount creates new account, deposit it with provided amount and returns it's hex private key. For testing purpose.
func NewFundedAccount(t *testing.T, pk, dialUri string, amount *big.Int) *ecdsa.PrivateKey {
	t.Helper()
	var (
		ctx        = context.Background()
		maxRetries = 10
		retryDelay = time.Millisecond * 10
	)

	sourcePrivateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		t.Fatalf("failed to load private key: %s", err)
	}

	destPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("failed to generate private key: %s", err)
	}

	client, err := ethclient.Dial(dialUri)
	if err != nil {
		t.Fatalf("failed to connect to %s: %s", dialUri, err)
	}

	destPublicKey := destPrivateKey.Public().(*ecdsa.PublicKey)
	destAddress := crypto.PubkeyToAddress(*destPublicKey)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		t.Fatalf("failed to get network ID: %s", err)
	}
	signer := types.NewLondonSigner(chainID)

	for range maxRetries {
		if err := deposit(ctx, client, signer, destAddress, sourcePrivateKey, amount); err != nil {
			if errs.Is(err, core.ErrNonceTooLow) || errs.Is(err, txpool.ErrReplaceUnderpriced) {
				time.Sleep(retryDelay)
				continue
			}

			t.Fatalf("failed to deposit: %s", err)
		}

		return destPrivateKey
	}

	t.Fatalf("failed to deposit account %d times", maxRetries)
	return nil
}

// PrivateKeyToHex returns the private key in hex format without 0x prefix for SDK compatibility.
func PrivateKeyToHex(priv *ecdsa.PrivateKey) string {
	return hexutil.Encode(crypto.FromECDSA(priv))[2:]
}

func deposit(ctx context.Context, client *ethclient.Client, signer types.Signer, dest common.Address, pk *ecdsa.PrivateKey, amount *big.Int) error {
	sourcePublicKey := pk.Public().(*ecdsa.PublicKey)
	sourceAddress := crypto.PubkeyToAddress(*sourcePublicKey)

	mu.Lock()
	defer mu.Unlock()
	nonce, err := client.PendingNonceAt(ctx, sourceAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	gasLimit := uint64(21000)

	var tx *types.Transaction

	txData := &types.DynamicFeeTx{
		ChainID:   signer.ChainID(),
		Nonce:     nonce,
		GasTipCap: gasPrice,
		GasFeeCap: new(big.Int).Mul(gasPrice, big.NewInt(2)),
		Gas:       gasLimit,
		To:        &dest,
		Value:     amount,
		Data:      nil,
	}
	tx = types.NewTx(txData)

	signedTx, err := types.SignTx(tx, signer, pk)
	if err != nil {
		return err
	}

	if err := client.SendTransaction(ctx, signedTx); err != nil {
		return err
	}

	return WaitForTx(ctx, client, signedTx.Hash())
}

// ToWei convert amount into *big.Int.
func ToWei(amount int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(amount), big.NewInt(1e18))
}

// WaitForTx block execution until transaction receipt is received or context is cancelled.
func WaitForTx(ctx context.Context, ethClient *ethclient.Client, hash common.Hash) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	receipt, err := ethClient.TransactionReceipt(ctx, hash)
	if err == nil {
		if receipt.Status == 1 {
			return nil
		}

		return errs.New("transaction failed")
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
			receipt, err := ethClient.TransactionReceipt(ctx, hash)
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
