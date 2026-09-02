// Package cachestats - Gopher Workplace challenge.
package cachestats

import "sync/atomic"

// Stats counts cache hits and misses on a CDN edge node.
type Stats struct {
	hits   atomic.Int64
	misses atomic.Int64
}

// Hit records one cache hit.
//
// Examples:
//
//	var s Stats; s.Hit(); s.Hits() => 1
func (s *Stats) Hit() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Miss records one cache miss.
//
// Examples:
//
//	var s Stats; s.Miss(); s.Misses() => 1
func (s *Stats) Miss() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Hits returns the number of cache hits.
func (s *Stats) Hits() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Misses returns the number of cache misses.
func (s *Stats) Misses() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Ratio returns hits divided by total lookups, or 0 when there were none.
//
// Examples:
//
//	var s Stats; s.Hit(); s.Miss(); s.Ratio() => 0.5
//	var s Stats; s.Ratio()                    => 0
func (s *Stats) Ratio() float64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}
