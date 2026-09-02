# Graceful Shutdown Flag

## Intuition

`atomic.Bool` guarantees each read sees a whole, current value. `CompareAndSwap(false, true)` is a race that exactly one goroutine wins, which makes it the standard "claim this once" primitive - perfect for running the drain exactly once.

## Approach

1. Declare `down atomic.Bool`.
2. `Request` calls `Store(true)`; `Requested` calls `Load`.
3. `ClaimShutdown` returns `CompareAndSwap(false, true)`.

## Solution

```go
// Package shutdownflag - Gopher Workplace challenge.
package shutdownflag

import "sync/atomic"

// ShutdownFlag is a one-way graceful-shutdown flag safe for concurrent use.
type ShutdownFlag struct {
	down atomic.Bool
}

// Request marks shutdown as requested.
//
// Examples:
//
//	var f ShutdownFlag; f.Request(); f.Requested() => true
func (f *ShutdownFlag) Request() {
	f.down.Store(true)
}

// Requested reports whether shutdown has been requested.
//
// Examples:
//
//	var f ShutdownFlag; f.Requested() => false
func (f *ShutdownFlag) Requested() bool {
	return f.down.Load()
}

// ClaimShutdown requests shutdown and reports whether this call started it.
//
// Examples:
//
//	var f ShutdownFlag; f.ClaimShutdown() => true
//	f.ClaimShutdown()                     => false
func (f *ShutdownFlag) ClaimShutdown() bool {
	return f.down.CompareAndSwap(false, true)
}
```

## Walkthrough

A signal handler and a health check both call `ClaimShutdown`. The first CAS sees `false` and swaps to `true`, returning true - it runs the drain. The second sees `true`, fails, returns false and skips the drain.

## Pitfalls

- Using `if !f.Requested() { f.Request() }` - two goroutines can both pass the check and both drain.
- Storing into a plain `bool` field alongside atomic reads.
- Assuming `Store` reports whether the value changed; it returns nothing.
