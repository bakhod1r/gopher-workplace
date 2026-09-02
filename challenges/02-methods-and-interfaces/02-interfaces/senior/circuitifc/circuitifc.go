// Package circuitifc — Gopher Workplace challenge.
package circuitifc

import (
	"errors"
	"time"
)

// ErrOpen is returned while the breaker is open.
var ErrOpen = errors.New("circuit open")

// Clock reports the current time.
type Clock interface {
	Now() time.Time
}

// FakeClock is a manually advanced clock.
type FakeClock struct {
	T time.Time
}

// Now returns the current fake time.
func (f *FakeClock) Now() time.Time { return f.T }

// Advance moves the clock forward.
func (f *FakeClock) Advance(d time.Duration) { f.T = f.T.Add(d) }

// Op is the protected operation.
type Op interface {
	Do() error
}

// OpFunc adapts a function to Op.
type OpFunc func() error

// Do calls the underlying function.
func (f OpFunc) Do() error { return f() }

// Breaker guards an operation.
type Breaker struct {
	Threshold int
	Cooldown  time.Duration

	clock    Clock
	failures int
	open     bool
	openedAt time.Time
}

// NewBreaker returns a closed breaker.
func NewBreaker(threshold int, cooldown time.Duration, c Clock) *Breaker {
	return &Breaker{Threshold: threshold, Cooldown: cooldown, clock: c}
}

// Call runs op unless the breaker is open.
//
// Examples:
//
//	closed     => op runs, its error is returned
//	open       => ErrOpen, op is not called
//	after cooldown => one probe is allowed through
func (b *Breaker) Call(op Op) error {
	// TODO(candidate): fail fast while open; probe after the cooldown.
	panic("not implemented")
}

// IsOpen reports the breaker state.
func (b *Breaker) IsOpen() bool {
	// TODO(candidate): report the state.
	panic("not implemented")
}
