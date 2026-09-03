# Scheduler Leader Lease

## Intuition

Every replica shouts "I am leader of term 5!" at the same moment. `CompareAndSwap` is the referee: only the caller who finds the term still at 4 gets to write 5, and the hardware guarantees exactly one of them does.

## Approach

1. `Claim` returns `l.term.CompareAndSwap(term-1, term)`.
2. `Term` returns `l.term.Load()`.

## Solution

```go
func (l *Lease) Claim(term int64) bool {
	return l.term.CompareAndSwap(term-1, term)
}

func (l *Lease) Term() int64 {
	return l.term.Load()
}
```

## Walkthrough

A fresh `Lease` holds 0. `Claim(1)` compares 0 with the expected 0, matches, stores 1 and returns true. A second `Claim(1)` compares 1 with 0, fails, returns false. `Claim(5)` on a lease at 0 expects 4 — no match, false.

## Pitfalls

- Reading with `Load` then storing with `Store` is two steps; two replicas can both win.
- Guarding an ordinary `int64` with a mutex works but is heavier than a single CAS.
- Forgetting that the zero term is already "taken" — `Claim(0)` expects -1 and must fail.
