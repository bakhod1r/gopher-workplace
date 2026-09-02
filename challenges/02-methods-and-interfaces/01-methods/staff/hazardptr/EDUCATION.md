# Hazard Pointer

## Intuition

In a lock-free structure, "is anyone still reading this node?" has no cheap
answer — so hazard pointers turn the question around. Each reader publishes the
single address it is about to dereference. A thread that wants to free memory
first scans everyone's hazard slots and skips anything listed.

The subtlety is entirely in the ordering. Announcing *after* validating would
let a reclaimer scan the empty slot, free the object, and leave the reader
holding a dangling pointer it just declared safe.

## Approach

1. Load the current pointer.
2. Publish it as this thread's hazard.
3. Re-load and compare — this proves the announcement was visible before the
   pointer could be retired.
4. On mismatch, give up and let the caller retry.

## Solution

```go
func (h *Hazard) Protect(shared *atomic.Pointer[int]) *int {
	p := shared.Load()
	h.ptr.Store(p)
	if shared.Load() == p {
		return p
	}
	return nil
}
```

## Walkthrough

The test stores `&val` into `shared` and calls `Protect`. Nothing else is
racing, so both loads see the same address, the comparison holds, and the caller
gets a pointer that dereferences to 42.

Under contention the second load can differ: a writer swapped the pointer in
between. The announcement is then stale, so `Protect` returns nil rather than a
pointer that may already be scheduled for reclamation.

## Pitfalls

- **Storing the hazard after the validation.** The classic mis-ordering; it
  reopens the exact window the protocol closes.
- **Skipping the second load.** Without it the announcement proves nothing —
  the pointer could have been retired between the load and the store.
- **Returning `p` on mismatch.** Hands back a pointer that may be freed.
- **Not clearing `h.ptr` when done.** Real implementations must release the
  hazard, or the object is pinned forever.

## Why `atomic.Pointer[int]` and not a plain `*int`

A non-atomic pointer read racing with a write is undefined behaviour in Go's
memory model — the race detector flags it, and the compiler is free to reload or
cache the value. `atomic.Pointer[T]`, added in Go 1.19, gives a typed,
sequentially consistent load and store without any `unsafe` conversions.
