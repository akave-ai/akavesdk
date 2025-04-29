// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package merkletree_test

import (
	"crypto/rand"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/merkletree"
)

func TestPaddedReader(t *testing.T) {
	size := 12 * memory.MB
	paddedSize := 16777216 * memory.B

	// example random file data.
	dataR := io.LimitReader(rand.Reader, size.ToInt64())
	paddedR := merkletree.NewPaddedReader(dataR, paddedSize.ToInt64())

	f, err := os.CreateTemp(t.TempDir(), "")
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, f.Close())
	})

	n, err := io.Copy(f, paddedR)
	require.NoError(t, err)
	stat, err := f.Stat()
	require.NoError(t, err)

	require.Equal(t, paddedSize.ToInt64(), n)
	require.Equal(t, paddedSize.ToInt64(), stat.Size())

	padded := paddedSize - size
	buf := make([]byte, int(padded))

	at, err := f.ReadAt(buf, size.ToInt64())
	require.NoError(t, err)
	require.Equal(t, int(padded), at)
	require.EqualValues(t, make([]byte, int(padded)), buf)
}
