// Package chunksize — Gopher Workplace challenge.
package chunksize

import "errors"

// Chunking failures.
var (
	ErrBadParts      = errors.New("parts must be positive")
	ErrNegativeTotal = errors.New("total must not be negative")
)

// ChunkSize returns the smallest chunk size covering total in parts chunks.
//
// Examples:
//
//	ChunkSize(10, 3) => 4, nil
//	ChunkSize(10, 0) => 0, ErrBadParts
func ChunkSize(total, parts int) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
