// Package backoff — Gopher Workplace challenge.
package backoff

import "time"

// Backoff manages retry delays.
type Backoff struct {
	current time.Duration
	max     time.Duration
}

// New creates a Backoff starting at 1s, capped at max.
func New(max time.Duration) *Backoff {
	return &Backoff{current: time.Second, max: max}
}

// Next returns the current delay and doubles it for the next call,
// capped at max.
func (b *Backoff) Next() time.Duration {
	// TODO(candidate): return current, then double current (up to max).
	panic("not implemented")
}
