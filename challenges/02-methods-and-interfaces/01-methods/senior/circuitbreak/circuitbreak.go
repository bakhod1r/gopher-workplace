// Package circuitbreak — Gopher Workplace challenge.
package circuitbreak

import "errors"

// Breaker implements a simple circuit breaker.
type Breaker struct {
	ConsecutiveFails int
	Threshold        int
	IsOpen           bool
}

// Call executes fn. If it returns an error, fails are incremented.
// If fails >= Threshold, IsOpen becomes true.
// If IsOpen is true, Call returns an error immediately without executing fn.
// If fn succeeds, fails are reset to 0.
func (b *Breaker) Call(fn func() error) error {
	// TODO(candidate): implement circuit breaker logic.
	_ = errors.New
	panic("not implemented")
}
