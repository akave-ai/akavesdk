// Copyright (C) 2025 Akave
// See LICENSE for copying information.

// Package ipctest provides an ipc testing utils.
package ipctest

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zeebo/errs"
)

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

	return client.SendTransaction(ctx, signedTx)
}

// ToWei convert amount into *big.Int.
func ToWei(amount int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(amount), big.NewInt(1e18))
}
