// Package endpointhits - Gopher Workplace challenge.
package endpointhits

import "sync"

// HitCounter counts gateway requests per route.
type HitCounter struct {
	mu   sync.Mutex
	hits map[string]int
}

// NewHitCounter returns an empty per-route counter.
func NewHitCounter() *HitCounter {
	return &HitCounter{hits: make(map[string]int)}
}

// Record counts one request served by route.
//
// Examples:
//
//	h.Record("/users"); h.Hits("/users")                     => 1
//	h.Record("/users"); h.Record("/users"); h.Hits("/users") => 2
func (h *HitCounter) Record(route string) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Hits returns the number of requests served by route.
//
// Examples:
//
//	NewHitCounter().Hits("/orders") => 0
func (h *HitCounter) Hits(route string) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
