# One-shot execution via captured state

## Intuition

A captured boolean lets a closure enforce run-once semantics across many calls.

## Approach

1. Capture a `done` flag.
2. Run `f` only when `!done`, then set `done = true`.

## Solution

```go
func Once(f func()) func() {
	done := false
	return func() {
		if !done {
			done = true
			f()
		}
	}
}
```

## Walkthrough

The first call runs `f` and flips `done`; every later call sees `done` true and does nothing.

## Pitfalls

- Set the flag BEFORE or after calling f consistently; here order doesn't matter (single-threaded).
- Not goroutine-safe; that's `sync.Once`'s job.
