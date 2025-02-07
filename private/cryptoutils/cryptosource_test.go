// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package cryptoutils_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/cryptoutils"
)

func TestCryptosource(t *testing.T) {
	rand := rand.New(cryptoutils.NewSource())
	for range 100 {
		require.Less(t, rand.Intn(100), 100)
	}
}

func TestShuffle(t *testing.T) {
	peers := []string{"a", "b", "c", "d", "e"}
	rand := rand.New(cryptoutils.NewSource())

	require.EventuallyWithT(t, func(collect *assert.CollectT) {
		rand.Shuffle(len(peers), func(i, j int) {
			peers[i], peers[j] = peers[j], peers[i]
		})
		assert.ElementsMatch(collect, []string{"a", "b", "c", "d", "e"}, peers)
		assert.NotEqual(collect, []string{"a", "b", "c", "d", "e"}, peers)
	}, 1*time.Second, 200*time.Millisecond)
}
