// Copyright (C) 2026 Akave
// See LICENSE for copying information.

package erasurecode

import "github.com/akave-ai/akavesdk/private/encryption"

// SplitStripes splits data into stripes of the given maxStripeSize. The last stripe may be smaller.
func SplitStripes(data []byte, maxStripeSize int) [][]byte {
	n := encryption.CeilDiv(len(data), maxStripeSize)
	result := make([][]byte, n)
	for i := range result {
		start := i * maxStripeSize
		end := min(start+maxStripeSize, len(data))
		result[i] = data[start:end:end]
	}
	return result
}
