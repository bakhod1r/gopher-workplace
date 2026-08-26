// Package syncpool — Gopher Workplace challenge.
package syncpool

import "sync"

// Buffer is a dummy struct we want to pool.
type Buffer struct {
	Data []byte
}

// BufferPool wraps sync.Pool to provide type-safety.
type BufferPool struct {
	pool sync.Pool
}

// NewPool creates a pool.
func NewPool() *BufferPool {
	return &BufferPool{
		pool: sync.Pool{
			New: func() interface{} { return &Buffer{Data: make([]byte, 1024)} },
		},
	}
}

// Get returns a Buffer from the pool.
func (p *BufferPool) Get() *Buffer {
	// TODO(candidate): get from p.pool, assert to *Buffer
	panic("not implemented")
}

// Put returns a Buffer to the pool.
func (p *BufferPool) Put(b *Buffer) {
	// TODO(candidate): put b into p.pool
	panic("not implemented")
}
