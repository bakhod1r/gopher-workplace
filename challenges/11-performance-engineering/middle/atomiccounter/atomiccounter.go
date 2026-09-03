// Package atomiccounter — Gopher Workplace challenge.
package atomiccounter

import "sync/atomic"

// Stats tracks a total and a high-water mark without a mutex. The zero value
// is ready to use and safe for concurrent use.
type Stats struct {
	total atomic.Int64
	max   atomic.Int64
}

// Add increases the total by delta and returns the new value.
//
// Examples:
//
//	s.Add(3) => 3
func (s *Stats) Add(delta int64) int64 {
	panic("not implemented")
}

// Total returns the current total.
//
// Examples:
//
//	s.Total() => 3
func (s *Stats) Total() int64 {
	panic("not implemented")
}

// Observe records a value and keeps the largest seen so far. A load, a
// comparison and a store is not atomic, so this needs a compare-and-swap
// retry loop: read the current max, and only store if it is still what you
// read.
//
// Examples:
//
//	s.Observe(5); s.Observe(2); s.Max() => 5
func (s *Stats) Observe(v int64) {
	panic("not implemented")
}

// Max returns the largest value observed, or 0 if none.
//
// Examples:
//
//	s.Max() => 5
func (s *Stats) Max() int64 {
	panic("not implemented")
}
