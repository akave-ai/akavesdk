// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package ipc_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	commcid "github.com/filecoin-project/go-fil-commcid"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/stretchr/testify/require"
	"github.com/zeebo/errs"

	"github.com/akave-ai/akavesdk/private/ipc/contracts"
	"github.com/akave-ai/akavesdk/private/ipctest"
	"github.com/akave-ai/akavesdk/private/merkletree"
)

func TestContractPDP(t *testing.T) {
	var (
		ctx        = context.Background()
		dialUri    = PickDialURI(t)
		privateKey = PickPrivateKey(t)
	)

	pk := ipctest.NewFundedAccount(t, privateKey, dialUri, ipctest.ToWei(10))

	ethClient, err := ethclient.Dial(dialUri)
	require.NoError(t, err)

	dialUriWs := strings.Replace(dialUri, "http", "ws", 1)
	wsClient, err := ethclient.Dial(dialUriWs)
	require.NoError(t, err)

	chainID, err := ethClient.ChainID(ctx)
	require.NoError(t, err)

	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	require.NoError(t, err)

	authWithValue := *auth
	authWithValue.Value = big.NewInt(0).SetInt64(1e18)

	challengeFinality := 3

	verifierAddress, tx, pdpVerifier, err := contracts.DeployPDPVerifier(auth, ethClient, big.NewInt(int64(challengeFinality)))
	require.NoError(t, err)
	require.NoError(t, waitForTx(ethClient, tx.Hash()))

	sinkAddress, tx, _, err := contracts.DeploySink(auth, ethClient)
	require.NoError(t, err)
	require.NoError(t, waitForTx(ethClient, tx.Hash()))

	pdpContract, err := contracts.NewPDPVerifierFilterer(verifierAddress, wsClient)
	require.NoError(t, err)

	createProofSetCh := make(chan *contracts.PDPVerifierProofSetCreated)
	setId := big.NewInt(0)

	sub, err := pdpContract.WatchProofSetCreated(
		&bind.WatchOpts{Context: ctx},
		createProofSetCh,
		nil,
		[]common.Address{auth.From},
	)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	tx, err = pdpVerifier.CreateProofSet(&authWithValue, sinkAddress, nil)
	require.NoError(t, err)
	require.NoError(t, waitForTx(ethClient, tx.Hash()))

	select {
	case err := <-sub.Err():
		require.NoError(t, err)
	case event := <-createProofSetCh:
		setId.SetInt64(event.SetId.Int64())
	case <-ctx.Done():
		t.Fatal(ctx.Err())
	}
	sub.Unsubscribe()
	close(createProofSetCh)

	cp := new(commp.Calc)
	data := []byte{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
		1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
		2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
		3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
		4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
		5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
		6, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
		7, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
		8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	}

	_, err = cp.Write(data)
	require.NoError(t, err)

	rawCommP, paddedSize, err := cp.Digest()
	require.NoError(t, err)

	commCid, err := commcid.DataCommitmentV1ToCID(rawCommP)
	require.NoError(t, err)

	tree, err := merkletree.BuildSha254Memtree(
		merkletree.NewPaddedReader(bytes.NewReader(data), int64(paddedSize)),
		abi.PaddedPieceSize(paddedSize).Unpadded())
	require.NoError(t, err)
	require.NotNil(t, tree)

	rootData := []contracts.PDPVerifierRootData{{
		Root:    contracts.CidsCid{Data: commCid.Bytes()},
		RawSize: big.NewInt(0).SetUint64(paddedSize),
	}}

	tx, err = pdpVerifier.AddRoots(auth, setId, rootData, []byte{})
	require.NoError(t, err)
	require.NoError(t, waitForTx(ethClient, tx.Hash()))

	block, err := ethClient.BlockNumber(ctx)
	require.NoError(t, err)

	tx, err = pdpVerifier.NextProvingPeriod(auth, setId, big.NewInt(int64(block)+4), nil)
	require.NoError(t, err)
	require.NoError(t, waitForTx(ethClient, tx.Hash()))

	for i := 0; i < 4; i++ {
		fillBlocks(t)
	}

	epoch, err := pdpVerifier.GetNextChallengeEpoch(&bind.CallOpts{Context: ctx}, setId)
	require.NoError(t, err)

	var proofsToProve []contracts.PDPVerifierProof
	for i := 0; i < 6; i++ {
		epochBytes := padLeft(epoch.Bytes(), 32)
		setIdBytes := padLeft(setId.Bytes(), 32)
		iBytes := padLeft(big.NewInt(int64(i)).Bytes(), 8)

		var payload []byte
		payload = append(payload, epochBytes...)
		payload = append(payload, setIdBytes...)
		payload = append(payload, iBytes...)

		hash := crypto.Keccak256(payload)

		hashInt := new(big.Int).SetBytes(hash)

		leafCount := int64(paddedSize / 32)
		challengeIdx := new(big.Int).Mod(hashInt, big.NewInt(leafCount))

		proof, err := merkletree.MemtreeProof(tree, challengeIdx.Int64())
		require.NoError(t, err)

		proofsToProve = append(proofsToProve, contracts.PDPVerifierProof{
			Leaf:  proof.Leaf,
			Proof: proof.Proof,
		})
	}

	tx, err = pdpVerifier.ProvePossession(&authWithValue, setId, proofsToProve)
	require.NoError(t, err)
	require.NoError(t, waitForTx(ethClient, tx.Hash()))
}

func fillBlocks(t *testing.T) {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "anvil_mine",
		"params":  []interface{}{},
		"id":      1,
	}
	body, err := json.Marshal(payload)
	require.NoError(t, err)

	req, err := http.NewRequestWithContext(context.Background(), "POST", "http://127.0.0.1:8545", bytes.NewBuffer(body))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer require.NoError(t, resp.Body.Close())
}

func waitForTx(client *ethclient.Client, hash common.Hash) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ticker := time.NewTicker(200 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			receipt, err := client.TransactionReceipt(ctx, hash)
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

// Helper function to pad bytes on the left with zeros.
func padLeft(bytes []byte, size int) []byte {
	if len(bytes) >= size {
		return bytes
	}

	paddedBytes := make([]byte, size)
	copy(paddedBytes[size-len(bytes):], bytes)

	return paddedBytes
}
