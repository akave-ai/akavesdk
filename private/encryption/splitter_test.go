// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package encryption_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/encryption"
)

func TestSplitter(t *testing.T) {
	secretKey := []byte("N1PCdw3M2B1TfJhoaY2mL736p2vCUc47")
	data := "big brown fox jumps over the lazy dog\n\nfoo"

	splitter, err := encryption.NewSplitter(secretKey, bytes.NewReader([]byte(data)), 5)
	require.NoError(t, err)

	splitted := make([][]byte, 0)

	temp, err := splitter.NextBytes()
	require.NoError(t, err)
	splitted = append(splitted, temp)

	// continiuosly read the data from the splitter
	for err == nil {
		temp, err = splitter.NextBytes()

		require.Condition(t, func() bool {
			return errors.Is(err, io.EOF) || err == nil
		})

		if err == nil {
			splitted = append(splitted, temp)
		}
	}

	// decrypt data
	joinedData := make([]byte, 0)
	for i, data := range splitted {
		infoString := fmt.Sprintf("block_%d", i)
		decrypted, err := encryption.Decrypt(secretKey, data, []byte(infoString))
		require.NoError(t, err)
		joinedData = append(joinedData, decrypted...)
	}

	require.Equal(t, data, string(joinedData))
}
