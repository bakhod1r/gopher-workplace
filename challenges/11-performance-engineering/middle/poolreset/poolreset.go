// Package poolreset — Gopher Workplace challenge.
package poolreset

import "sync"

// MaxCap is the largest buffer this pool will keep. Anything bigger is
// dropped on Put so one huge request cannot pin a huge buffer forever.
const MaxCap = 64 * 1024

// Pool recycles byte buffers with a size ceiling.
type Pool struct {
	pool sync.Pool
	kept int // number of buffers accepted by Put, for the tests
	mu   sync.Mutex
}

// Put returns a buffer for reuse and reports whether the pool kept it.
// Buffers with capacity above MaxCap, and nil buffers, are dropped: a pool
// that keeps everything grows to the largest request ever seen and never
// gives that memory back.
//
// Examples:
//
//	p.Put(make([]byte, 0, 1024)) => true
//	p.Put(make([]byte, 0, 1<<20)) => false
func (p *Pool) Put(b []byte) bool {
	panic("not implemented")
}

// Get returns an empty buffer with capacity at least n, allocating when the
// pool has nothing suitable. A non-positive n gives a buffer with the pool's
// default capacity of 1024.
//
// Examples:
//
//	b := p.Get(4096); len(b) == 0 && cap(b) >= 4096
func (p *Pool) Get(n int) []byte {
	panic("not implemented")
}

// Kept reports how many buffers Put has accepted.
//
// Examples:
//
//	p.Kept() => 1
func (p *Pool) Kept() int {
	panic("not implemented")
}
