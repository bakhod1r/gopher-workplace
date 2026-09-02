# Lock Escalation

## Intuition

Double-checked locking exists because acquiring the expensive lock is worth
avoiding, but the cheap check is not authoritative. So you check cheaply, and if
it looks like work is needed, you take the real lock and check again — this time
under a lock that can actually be trusted.

In Go the extra wrinkle is that `sync.RWMutex` has no upgrade path. There is no
"promote my read lock"; you release and re-acquire, and you must assume the
world changed in between.

## Approach

1. Fast path under `RLock`: if there is nothing to do, return.
2. Release the read lock completely.
3. Take the write lock and re-test the condition.
4. Mutate if it still holds, capture the value, unlock.

## Solution

```go
func (o *OptLock) IncrementIfZero() int {
	o.mu.RLock()
	if o.v != 0 {
		v := o.v
		o.mu.RUnlock()
		return v
	}
	o.mu.RUnlock()

	o.mu.Lock()
	if o.v == 0 {
		o.v++
	}
	v := o.v
	o.mu.Unlock()
	return v
}
```

## Walkthrough

First call: the read shows 0, so the fast path does not fire. The read lock is
released, the write lock taken, the re-check still sees 0, and `v` becomes 1.

Second call: the read shows 1, so it returns immediately without ever touching
the write lock — that is the optimization.

Under contention, two goroutines can both pass the first check. One takes the
write lock and increments; the other then acquires it, finds `o.v == 1`, and
skips. Without that second check both would increment and the value would be 2.

## Pitfalls

- **Calling `Lock` while holding `RLock`.** `sync.RWMutex` is not reentrant or
  upgradable; this deadlocks the goroutine against itself.
- **Skipping the re-check.** The double-increment race — invisible in a
  single-goroutine test, which is why this one is worth reasoning about rather
  than running.
- **`defer o.mu.RUnlock()` at the top.** The deferred unlock runs at *return*,
  so the read lock is still held when `Lock` is called. Same deadlock.
- **Reading `o.v` after `Unlock`.** An unsynchronized read; the race detector
  flags it.

## Is the fast path worth it?

Only when reads dominate and the write is rare. `RLock` is not free — it still
does an atomic add on a shared cache line — so for a single int, `sync/atomic`
or a plain `Mutex` is often faster. The pattern earns its keep when the guarded
work is substantial, as in lazy initialization of an expensive value.
