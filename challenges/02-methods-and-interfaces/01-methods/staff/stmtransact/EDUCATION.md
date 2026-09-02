# Optimistic Transaction

## Intuition

Pessimistic concurrency takes a lock and hopes it is short. Optimistic
concurrency takes no lock at all, does the work, and asks at the last moment
whether the world moved. When conflicts are rare that is much cheaper; when they
are common it degenerates into repeated retries — which is the trade-off, not a
bug.

The whole correctness argument rests on one thing: a retry must recompute, not
resubmit.

## Approach

1. Read value and version together.
2. Compute the new value outside any lock.
3. Attempt a versioned commit.
4. Retry the whole thing on rejection.

## Solution

```go
func (tv *TVar) Tx(fn func(int) int) int {
	for {
		val, version := tv.Read()
		next := fn(val)
		if tv.Commit(version, next) {
			return next
		}
	}
}
```

## Walkthrough

Two goroutines both read `(0, version 0)` and both compute 1. The first
`Commit(0, 1)` succeeds and bumps the version to 1. The second `Commit(0, 1)`
finds `tv.version == 1 != 0` and is rejected. The loop reads again — now
`(1, version 1)` — recomputes `fn(1) == 2`, and commits.

That recomputation is why the 100-goroutine test lands on exactly 100. Move
`fn` above the loop and the retry commits a stale `1` over and over: the final
value comes out well under 100, and only under contention, which is exactly the
kind of bug that never reproduces locally.

## Pitfalls

- **Hoisting `fn(val)` out of the loop.** Lost updates under contention; the
  single-threaded tests still pass.
- **Comparing values instead of versions.** The ABA problem: a variable changed
  from 5 to 7 and back to 5 looks untouched, and a transaction that should have
  been rejected commits.
- **Holding a lock across `fn`.** Correct, but that is pessimistic locking with
  extra steps — and a slow or blocking `fn` now stalls every other writer.
- **Bounding the retries.** Sometimes wanted, but then `Tx` must be able to
  report failure; the signature here promises eventual success.

## Not the same as a lock-free algorithm

This loop can starve an individual transaction — under permanent contention one
goroutine may retry indefinitely — so it is lock-free (the system as a whole
progresses) but not wait-free (no per-operation bound). Real STM implementations
add contention managers and read sets spanning many variables.
