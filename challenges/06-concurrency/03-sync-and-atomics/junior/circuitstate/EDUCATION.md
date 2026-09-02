# Circuit Breaker State

## Intuition

A state machine transition is a conditional write: change to `open` only if the state is still `closed`. CAS expresses exactly that, and its boolean result names the single goroutine that performed the transition — the one that should run the side effects.

## Approach

1. Represent the state as an `atomic.Int32` with constants `closed = 0`, `open = 1`.
2. `Trip` returns `CompareAndSwap(closed, open)`.
3. `Reset` returns `CompareAndSwap(open, closed)`; `Open` compares `Load()` to `open`.

## Solution

```go
// Package circuitstate - Gopher Workplace challenge.
package circuitstate

import "sync/atomic"

const (
	closed int32 = 0
	opened int32 = 1
)

// Breaker is a two-state circuit breaker safe for concurrent use.
type Breaker struct {
	state atomic.Int32
}

// Trip opens the breaker and reports whether this call opened it.
//
// Examples:
//
//	var b Breaker; b.Trip() => true
//	b.Trip()                => false
func (b *Breaker) Trip() bool {
	return b.state.CompareAndSwap(closed, opened)
}

// Reset closes the breaker and reports whether this call closed it.
//
// Examples:
//
//	var b Breaker; b.Trip(); b.Reset() => true
//	var b Breaker; b.Reset()           => false
func (b *Breaker) Reset() bool {
	return b.state.CompareAndSwap(opened, closed)
}

// Open reports whether the breaker is currently open.
//
// Examples:
//
//	var b Breaker; b.Open() => false
func (b *Breaker) Open() bool {
	return b.state.Load() == opened
}
```

## Walkthrough

Ten goroutines see a failure and all call `Trip`. One CAS finds `closed` and swaps to `open`, returning true — it logs the outage. The other nine find `open`, fail, and return false.

## Pitfalls

- `if !b.Open() { b.state.Store(open) }` — several goroutines pass the check and all think they tripped it.
- Reusing `Trip` to re-open an already-open breaker and expecting true.
- Reading the state field directly instead of through `Load`.
