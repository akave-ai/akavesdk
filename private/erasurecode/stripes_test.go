// Copyright (C) 2026 Akave
// See LICENSE for copying information.

package erasurecode_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/erasurecode"
	"github.com/akave-ai/akavesdk/private/testrand"
)

func TestSplitStripesErasureCodeEnabled(t *testing.T) {
	const stripeSize = 2 * 1024 * 1024 // 2 MiB

	data := testrand.Bytes(t, int64(4*stripeSize))
	result := erasurecode.SplitStripes(data, stripeSize)
	require.Len(t, result, 4)
	for _, s := range result {
		require.Len(t, s, stripeSize)
	}
}

func TestSplitStripesErasureCodeDisabled(t *testing.T) {
	const stripeSize = 2 * 1024 * 1024 // 2 MiB

	data := testrand.Bytes(t, int64(4*stripeSize))
	result := erasurecode.SplitStripes(data, 2*stripeSize)
	require.Len(t, result, 2)
	for _, s := range result {
		require.Len(t, s, 2*stripeSize)
	}
}

func TestSplitStripesLastStripeSmaller(t *testing.T) {
	const stripeSize = 2 * 1024 * 1024 // 2 MiB

	data := testrand.Bytes(t, int64(stripeSize+100))
	result := erasurecode.SplitStripes(data, stripeSize)
	require.Len(t, result, 2)
	require.Len(t, result[0], stripeSize)
	require.Len(t, result[1], 100)
}

func TestSplitStripesPreservesData(t *testing.T) {
	data := testrand.Bytes(t, 5*1024*1024)
	result := erasurecode.SplitStripes(data, 2*1024*1024)

	var gathered []byte
	for _, s := range result {
		gathered = append(gathered, s...)
	}
	require.Equal(t, data, gathered)
}

func TestSplitStripesEmpty(t *testing.T) {
	result := erasurecode.SplitStripes([]byte{}, 2*1024*1024)
	require.Empty(t, result)
}
