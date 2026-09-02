# Webhook Dedupe Filter

## Intuition

Separate `Load` then `Store` calls leave a gap where two deliveries both see "absent" and both get processed. `LoadOrStore` collapses both into one atomic operation, so exactly one caller gets `loaded == false`.

## Approach

1. Hold a `sync.Map` of event IDs.
2. `Accept` uses `LoadOrStore` and returns `!loaded`.
3. `Seen` uses `Load`; `Len` counts with `Range`.

## Solution

```go
// Package dedupefilter - Gopher Workplace challenge.
package dedupefilter

import "sync"

// DedupeFilter accepts each webhook event ID exactly once.
type DedupeFilter struct {
	seen sync.Map
}

// Accept reports whether this delivery is the first for eventID.
//
// Examples:
//
//	var d DedupeFilter; d.Accept("evt-1") => true
//	d.Accept("evt-1")                     => false
func (d *DedupeFilter) Accept(eventID string) bool {
	_, loaded := d.seen.LoadOrStore(eventID, struct{}{})
	return !loaded
}

// Seen reports whether eventID has ever been accepted.
//
// Examples:
//
//	var d DedupeFilter; d.Seen("evt-1") => false
func (d *DedupeFilter) Seen(eventID string) bool {
	_, ok := d.seen.Load(eventID)
	return ok
}

// Len returns the number of distinct accepted events.
//
// Examples:
//
//	d.Accept("evt-1"); d.Accept("evt-2"); d.Len() => 2
func (d *DedupeFilter) Len() int {
	n := 0
	d.seen.Range(func(_, _ any) bool {
		n++
		return true
	})
	return n
}
```

## Walkthrough

Two connections deliver `evt-1` at once. `LoadOrStore` serialises them: one stores and gets `loaded == false` (process it), the other reads the existing entry and gets `loaded == true` (drop it).

## Pitfalls

- Using `Load` then `Store` - both deliveries can slip through the gap.
- Expecting `sync.Map` to have a `Len` method; it does not, you must `Range`.
- Reaching for `sync.Map` by default; a plain map plus a mutex is usually faster.
