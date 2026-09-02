// Package circuitstate - Gopher Workplace challenge.
package circuitstate

import "sync/atomic"

const (
	closed int32 = 0
	opened int32 = 1
)

// Breaker is a two-state circuit breaker safe for concurrent use.
type Breaker struct {
	state atomic.Int32
}

// Trip opens the breaker and reports whether this call opened it.
//
// Examples:
//
//	var b Breaker; b.Trip() => true
//	b.Trip()                => false
func (b *Breaker) Trip() bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Reset closes the breaker and reports whether this call closed it.
//
// Examples:
//
//	var b Breaker; b.Trip(); b.Reset() => true
//	var b Breaker; b.Reset()           => false
func (b *Breaker) Reset() bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Open reports whether the breaker is currently open.
//
// Examples:
//
//	var b Breaker; b.Open() => false
func (b *Breaker) Open() bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
