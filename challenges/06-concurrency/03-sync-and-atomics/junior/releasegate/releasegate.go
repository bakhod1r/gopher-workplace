// Package releasegate - Gopher Workplace challenge.
package releasegate

import "sync"

// Gate holds goroutines until a release opens it.
type Gate struct {
	mu   sync.Mutex
	cond *sync.Cond
	open bool
}

// NewGate returns a closed gate.
func NewGate() *Gate {
	g := &Gate{}
	g.cond = sync.NewCond(&g.mu)
	return g
}

// Wait blocks until the gate is open.
//
// Examples:
//
//	g := NewGate(); g.Open(); g.Wait() => returns immediately
func (g *Gate) Wait() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Open opens the gate and releases every waiting goroutine.
//
// Examples:
//
//	g := NewGate(); g.Open(); g.IsOpen() => true
func (g *Gate) Open() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// IsOpen reports whether the gate is open.
//
// Examples:
//
//	g := NewGate(); g.IsOpen() => false
func (g *Gate) IsOpen() bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
