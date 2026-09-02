# Cache Hit Ratio

## Intuition

Each counter is safe on its own, but reading hits and misses takes two loads — traffic can arrive between them. For a monitoring ratio that skew is acceptable; load each counter exactly once into a local so at least the arithmetic is self-consistent.

## Approach

1. `Hit` and `Miss` call `Add(1)` on their counter.
2. `Ratio` loads both into locals.
3. If the total is 0 return 0, else return `float64(h) / float64(h+m)`.

## Solution

```go
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
	s.hits.Add(1)
}

// Miss records one cache miss.
//
// Examples:
//
//	var s Stats; s.Miss(); s.Misses() => 1
func (s *Stats) Miss() {
	s.misses.Add(1)
}

// Hits returns the number of cache hits.
func (s *Stats) Hits() int64 {
	return s.hits.Load()
}

// Misses returns the number of cache misses.
func (s *Stats) Misses() int64 {
	return s.misses.Load()
}

// Ratio returns hits divided by total lookups, or 0 when there were none.
//
// Examples:
//
//	var s Stats; s.Hit(); s.Miss(); s.Ratio() => 0.5
//	var s Stats; s.Ratio()                    => 0
func (s *Stats) Ratio() float64 {
	h := s.hits.Load()
	m := s.misses.Load()
	total := h + m
	if total == 0 {
		return 0
	}
	return float64(h) / float64(total)
}
```

## Walkthrough

Two hits and two misses are recorded from four goroutines. `Ratio` loads 2 and 2, the total is 4, and it returns `2.0 / 4.0` = 0.5.

## Pitfalls

- Loading a counter twice inside one expression, which can read two different values.
- Dividing without the zero guard, producing NaN for an idle node.
- Integer division: `h / (h + m)` on int64 truncates to 0.
