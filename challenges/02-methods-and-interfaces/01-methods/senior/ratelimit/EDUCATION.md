# Thread-Safe Methods

## Intuition

"Is there a token? Then take it." Those are two statements and one decision. If
another goroutine can run between them, two callers can both see the last token
and both take it — the bucket goes negative and the limit is exceeded. Holding
the mutex across both statements is what makes the pair atomic.

## Approach

1. Lock, and defer the unlock.
2. Test the count and decrement in the same critical section.
3. Refill under the same lock.

## Solution

```go
func (l *Limiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.tokens > 0 {
		l.tokens--
		return true
	}
	return false
}

func (l *Limiter) Refill(n int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.tokens += n
}
```

## Walkthrough

`NewLimiter(2)` starts with two tokens. Two `Allow` calls take them and return
true; the third finds 0, takes nothing and returns false. `Refill(1)` puts one
back, so the next `Allow` succeeds and the one after it fails again.

The mutex never protects `n` or the return value — only the shared `tokens`
field, which is exactly the scope a lock should have.

## Pitfalls

- **Decrementing before the check.** The count drifts negative while callers are
  refused, so a later `Refill(1)` does not actually admit anyone.
- **Locking only in `Allow`.** `Refill` then writes the same field unguarded — a
  data race the detector reports.
- **Returning before `Unlock` without `defer`.** The early `return false` path
  leaks the lock and every later call deadlocks.
- **Value receiver.** Copies the mutex and the counter; `go vet` flags the copy.

## What a real limiter adds

Time. A production token bucket refills continuously at a rate (tokens per
second) rather than on demand, usually by computing how many tokens *should*
have arrived since the last call. `golang.org/x/time/rate` does this without a
background goroutine — the same critical section, plus a clock reading.
