// Package ratecounter - Gopher Workplace challenge.
package ratecounter

import "sync"

// RateCounter counts gateway requests across many serving goroutines.
type RateCounter struct {
	mu   sync.Mutex
	hits int
}

// Record counts one inbound request.
//
// Examples:
//
//	var c RateCounter; c.Record(); c.Hits()             => 1
//	var c RateCounter; c.Record(); c.Record(); c.Hits() => 2
func (c *RateCounter) Record() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Hits returns the number of requests recorded so far.
//
// Examples:
//
//	var c RateCounter; c.Hits() => 0
func (c *RateCounter) Hits() int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
