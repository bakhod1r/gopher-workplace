// Package circuitstate — Gopher Workplace challenge.
package circuitstate

import "errors"

// ErrOpen reports a circuit that refuses further calls.
var ErrOpen = errors.New("circuit open")

// Breaker stops calling f after Threshold consecutive failures.
type Breaker struct {
	Threshold int

	failures int
}

// Call runs f unless the circuit is open.
//
// Examples:
//
//	(&Breaker{Threshold: 1}).Call(func() error { return nil }) => nil
func (b *Breaker) Call(f func() error) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
