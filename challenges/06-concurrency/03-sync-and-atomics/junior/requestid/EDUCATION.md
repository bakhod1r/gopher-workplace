# Request ID Generator

## Intuition

`Add` is one indivisible instruction that returns the post-increment value, so every caller receives a different number. Reading with `Load` and then storing `cur+1` opens a window where two handlers read the same `cur`.

## Approach

1. Declare `n atomic.Int64`.
2. `Next` returns `g.n.Add(1)`.
3. `Issued` returns `g.n.Load()`.

## Solution

```go
// Package requestid - Gopher Workplace challenge.
package requestid

import "sync/atomic"

// IDGen hands out unique, increasing request IDs.
type IDGen struct {
	n atomic.Int64
}

// Next returns the next unique request ID, starting at 1.
//
// Examples:
//
//	var g IDGen; g.Next()           => 1
//	var g IDGen; g.Next(); g.Next() => 2
func (g *IDGen) Next() int64 {
	return g.n.Add(1)
}

// Issued reports how many IDs have been handed out.
//
// Examples:
//
//	var g IDGen; g.Next(); g.Issued() => 1
func (g *IDGen) Issued() int64 {
	return g.n.Load()
}
```

## Walkthrough

Two handlers call `Next` on a fresh generator. The hardware serialises the two adds: one returns 1, the other 2. Neither can observe the other's intermediate state.

## Pitfalls

- `g.n.Store(g.n.Load() + 1)` — two separate operations, duplicate IDs.
- Returning `Load()` after `Add(1)`; another goroutine may have advanced it already.
- Starting from 0, which collides with the zero value of an unset ID field.
