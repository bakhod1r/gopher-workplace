// Package poolencoder — Gopher Workplace challenge.
package poolencoder

import "sync"

// Encoder renders fields into one line.
type Encoder interface {
	Encode(fields []string) string
}

// PooledEncoder reuses byte buffers across calls.
type PooledEncoder struct {
	pool sync.Pool
}

// NewPooledEncoder returns an encoder with a warm pool.
func NewPooledEncoder() *PooledEncoder {
	return &PooledEncoder{
		pool: sync.Pool{
			New: func() any {
				b := make([]byte, 0, 256)
				return &b
			},
		},
	}
}

// Encode joins the fields with "," reusing a pooled buffer.
//
// The returned string must not alias the pooled buffer.
//
// Examples:
//
//	Encode([]string{"a", "b"}) => "a,b"
func (e *PooledEncoder) Encode(fields []string) string {
	// TODO(candidate): take from the pool, reset, build, copy out, put back.
	panic("not implemented")
}

// EncodeAll encodes every message through enc.
func EncodeAll(enc Encoder, messages [][]string) []string {
	// TODO(candidate): one encoded line per message.
	panic("not implemented")
}
