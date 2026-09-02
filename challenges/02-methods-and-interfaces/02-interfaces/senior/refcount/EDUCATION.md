# Reference Counting

## Intuition

The dangerous window is between "the count reached zero" and "I closed it". Both must happen under one lock, or two goroutines can both observe zero and both close.

## Approach

1. `Acquire` refuses once released or at zero, otherwise increments.
2. `Release` refuses at zero, decrements, and closes when the count hits zero and `released` is still false.
3. Set `released` before calling `Close` so a re-entrant path cannot double-close.
4. `Count` reads under the same lock.

## Solution

```go
func (r *RefCounted) Acquire() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.released || r.refs <= 0 {
		return false
	}
	r.refs++
	return true
}

func (r *RefCounted) Release() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.refs <= 0 {
		return false
	}
	r.refs--
	if r.refs == 0 && !r.released {
		r.released = true
		r.res.Close()
	}
	return true
}
```

## Walkthrough

`TestManyHoldersReleaseConcurrently` fires 201 concurrent releases against 201 references. Exactly one of them observes the count reaching zero under the lock, so `ClosedTimes` is exactly 1.

## Pitfalls

- Using `atomic.AddInt32` for the count and closing when it returns 0 — correct here, but it cannot also guard `Acquire` against resurrection after close.
- Closing outside the lock, which opens the double-close window again.
- Allowing `Acquire` to revive a closed resource, handing out a dead handle.
