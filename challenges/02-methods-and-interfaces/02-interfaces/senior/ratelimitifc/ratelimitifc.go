// Package ratelimitifc — Gopher Workplace challenge.
package ratelimitifc

import "time"

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

// Limiter decides whether an action may proceed.
type Limiter interface {
	Allow() bool
}

// TokenBucket allows Burst actions, refilling one token per Interval.
type TokenBucket struct {
	Burst    int
	Interval time.Duration
	clock    Clock
	tokens   int
	last     time.Time
}

// NewTokenBucket returns a full bucket.
func NewTokenBucket(burst int, interval time.Duration, c Clock) *TokenBucket {
	return &TokenBucket{
		Burst:    burst,
		Interval: interval,
		clock:    c,
		tokens:   burst,
		last:     c.Now(),
	}
}

// Allow consumes a token if one is available.
//
// Examples:
//
//	Burst 2, no time passing => true, true, false
func (t *TokenBucket) Allow() bool {
	// TODO(candidate): refill by elapsed time (capped at Burst), then consume.
	panic("not implemented")
}

// AllowN reports how many of n attempts were allowed.
func (t *TokenBucket) AllowN(n int) int {
	// TODO(candidate): count the allowed attempts.
	panic("not implemented")
}
