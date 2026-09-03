# CAS Inventory Reservation

## Intuition

`Add(-n)` is atomic but unconditional: it will happily drive the count to -7. What you need is a *conditional* decrement, and CAS is how you build one. The loop body is 'here is what I saw, here is what I want; only apply it if nothing changed underneath me.' Contention makes the loop spin more, never makes it wrong.

## Approach

1. Reject `n <= 0` up front.
2. Enter an unbounded `for` loop.
3. `have := s.available.Load()` — a fresh read every iteration.
4. If `have < n`, there is genuinely not enough stock: return `false`.
5. `CompareAndSwap(have, have-n)`: on success return `true`, on failure loop again.
6. `Release`: ignore `n <= 0`, otherwise `s.available.Add(n)` — unconditional, so no loop needed.

## Solution

```go
// Reserve takes n units and reports whether the reservation succeeded. It
// fails when n is not positive or when fewer than n units remain. It is a
// compare-and-swap retry loop: read, decide, swap, and start over if another
// goroutine moved the count in between.
//
// Examples:
//
//	s := NewStock(10); s.Reserve(3)  => true
//	s := NewStock(2); s.Reserve(5)   => false
//	NewStock(10).Reserve(0)          => false
func (s *Stock) Reserve(n int64) bool {
	if n <= 0 {
		return false
	}
	for {
		have := s.available.Load()
		if have < n {
			return false
		}
		if s.available.CompareAndSwap(have, have-n) {
			return true
		}
	}
}

// Release returns n units to the pool. A non-positive n is ignored.
//
// Examples:
//
//	s := NewStock(1); s.Reserve(1); s.Release(1); s.Available() => 1
func (s *Stock) Release(n int64) {
	if n <= 0 {
		return
	}
	s.available.Add(n)
}
```

## Walkthrough

- Two workers both load `have == 1` and both want 1 unit. Both call `CompareAndSwap(1, 0)`; the hardware lets exactly one succeed. The loser loops, loads `0`, sees `0 < 1`, and returns `false`.
- `TestStockNoOversell` has 16 workers make 1600 attempts against 500 units and asserts exactly 500 succeed — a lost update would show up as 501 or as a negative `Available`.
- `Release` needs no CAS because adding is unconditional and `Add` is already atomic.
- `NewStock(-3)` stores nothing, so `Available` is `0` and every `Reserve` fails.

## Pitfalls

- Hoisting `have := s.available.Load()` above the loop — after a failed CAS you would retry forever against a stale value.
- Using `if s.available.Load() >= n { s.available.Add(-n) }`: the gap between the check and the add is exactly where overselling happens.
- Returning `false` when the CAS fails; a CAS failure means *retry*, not *out of stock*.
- Guarding a `Reserve`/`Release` pair with a mutex as well as a CAS loop — the mutex makes the CAS pointless and reintroduces the contention you were avoiding.
