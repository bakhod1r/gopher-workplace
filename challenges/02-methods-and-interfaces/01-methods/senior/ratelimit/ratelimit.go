// Package ratelimit — Gopher Workplace challenge.
package ratelimit

import "sync"

// Limiter implements a simple token bucket.
type Limiter struct {
	mu     sync.Mutex
	tokens int
}

// NewLimiter creates a limiter with max tokens.
func NewLimiter(max int) *Limiter {
	return &Limiter{tokens: max}
}

// Allow returns true if a token is available, decrementing the count.
func (l *Limiter) Allow() bool {
	// TODO(candidate): thread-safe decrement.
	panic("not implemented")
}

// Refill adds n tokens, up to some max (assume infinite max for this puzzle).
func (l *Limiter) Refill(n int) {
	// TODO(candidate): thread-safe refill.
	panic("not implemented")
}
