// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package merkletree

import (
	"io"

	"github.com/zeebo/errs"
)

// PaddedReader pads with zeroes if underlying reader is smaller than target total size set.
type PaddedReader struct {
	r     io.Reader
	total int64
	read  int64
}

// NewPaddedReader creates new instance of padded reader.
func NewPaddedReader(r io.Reader, total int64) *PaddedReader {
	return &PaddedReader{r: r, total: total}
}

// Read reads until total bytes are read. If underlying reader is exhausted before reading desired total byte length
// PaddedReader fills remaining space with zeroes.
func (reader *PaddedReader) Read(p []byte) (n int, err error) {
	if reader.read >= reader.total {
		return 0, io.EOF
	}

	remaining := reader.total - reader.read
	if int64(len(p)) > remaining {
		p = p[:remaining]
	}

	n, err = reader.r.Read(p)
	reader.read += int64(n)

	if errs.Is(err, io.EOF) {
		padding := len(p) - n
		for i := n; i < n+padding; i++ {
			p[i] = 0
		}
		reader.read += int64(padding)
		return n + padding, nil
	}

	return n, err
}
