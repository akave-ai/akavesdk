// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package memory provides a type for memory size.
package memory

import "fmt"

// Size is a type for memory size.
type Size int64

const (
	// B is a byte.
	B Size = 1 << (10 * iota)
	// KiB is a kibibyte.
	KiB
	// MiB is a mebibyte.
	MiB
	// GiB is a gibibyte.
	GiB
	// TiB is a tebibyte.
	TiB
	// PiB is a pebibyte.
	PiB
	// EiB is an exbibyte.
	EiB
	// KB is a kilobyte.
	KB Size = 1e3
	// MB is a megabyte.
	MB Size = 1e6
	// GB is a gigabyte.
	GB Size = 1e9
	// TB is a terabyte.
	TB Size = 1e12
	// PB is a petabyte.
	PB Size = 1e15
	// EB is an exabyte.
	EB Size = 1e18
)

// String is for convert Size to string.
func (s Size) String() string {
	return FormatSize(s)
}

// ToInt64 is for convert Size to int64.
func (s Size) ToInt64() int64 {
	return int64(s)
}

// MulInt64 is for multiply Size by an integer.
func (s Size) MulInt64(n int64) Size {
	return s * Size(n)
}

// DivInt64 is for divide Size by an integer.
func (s Size) DivInt64(n int64) Size {
	return s / Size(n)
}

// FormatSize is for format Size to string.
func FormatSize(size Size) string {
	switch {
	case size >= EB:
		return fmt.Sprintf("%.2fEB", float64(size)/float64(EB))
	case size >= PB:
		return fmt.Sprintf("%.2fPB", float64(size)/float64(PB))
	case size >= TB:
		return fmt.Sprintf("%.2fTB", float64(size)/float64(TB))
	case size >= GB:
		return fmt.Sprintf("%.2fGB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.2fMB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.2fKB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%dB", size)
	}
}

// FormatBytes is for format bytes to string.
func FormatBytes(bytes int64) string {
	switch {
	case bytes >= GB.ToInt64():
		return fmt.Sprintf("%d GB", bytes/GB.ToInt64())
	case bytes >= MB.ToInt64():
		return fmt.Sprintf("%d MB", bytes/MB.ToInt64())
	case bytes >= KB.ToInt64():
		return fmt.Sprintf("%d KB", bytes/KB.ToInt64())
	default:
		return fmt.Sprintf("%d Bytes", bytes)
	}
}
