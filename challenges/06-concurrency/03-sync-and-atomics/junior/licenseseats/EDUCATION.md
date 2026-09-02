# License Seat Pool

## Intuition

`Add(-1)` is atomic but unconditional: with one seat left, two logins both decrement and the count goes to -1. CAS lets you attach a condition to the write — "take a seat only if the count I checked is still the count" — so the bound holds.

## Approach

1. Read the free count.
2. If it is 0, refuse.
3. Otherwise CAS it down by one; on failure loop and re-read.

## Solution

```go
// Package licenseseats - Gopher Workplace challenge.
package licenseseats

import "sync/atomic"

// SeatPool hands out a bounded number of concurrent license seats.
type SeatPool struct {
	free atomic.Int64
}

// NewSeatPool returns a pool with n free seats.
func NewSeatPool(n int64) *SeatPool {
	p := &SeatPool{}
	p.free.Store(n)
	return p
}

// TryAcquire takes a seat and reports whether one was available.
//
// Examples:
//
//	p := NewSeatPool(2); p.TryAcquire()                  => true
//	p := NewSeatPool(1); p.TryAcquire(); p.TryAcquire()  => false
func (p *SeatPool) TryAcquire() bool {
	for {
		cur := p.free.Load()
		if cur == 0 {
			return false
		}
		if p.free.CompareAndSwap(cur, cur-1) {
			return true
		}
	}
}

// Release returns a seat to the pool.
//
// Examples:
//
//	p := NewSeatPool(1); p.TryAcquire(); p.Release(); p.Free() => 1
func (p *SeatPool) Release() {
	p.free.Add(1)
}

// Free reports how many seats remain.
func (p *SeatPool) Free() int64 {
	return p.free.Load()
}
```

## Walkthrough

Twenty logins hit a pool of eight. Each CAS that succeeds lowers the count by exactly one; the losers re-read and try again until the count reaches 0, after which every remaining login is refused. Exactly eight seats are handed out.

## Pitfalls

- Using `if p.Free() > 0 { p.free.Add(-1) }` — the check and the decrement are separate, so the pool over-allocates.
- Forgetting to re-read `cur` inside the loop.
- Letting `Release` push the count above the original size.
