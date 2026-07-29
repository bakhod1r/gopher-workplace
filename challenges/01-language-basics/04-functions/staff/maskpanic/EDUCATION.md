# Panics during deferred unwinding

## Intuition

A panic in a deferred call replaces the panic currently unwinding; the last panic wins, so cleanup code that panics erases the original error.

## Approach

1. A panic inside a deferred function replaces the original panic value.
2. Remove the `panic("cleanup-failure")` so the first panic survives.

## Solution

```go
func FirstPanic(f func()) (got any) {
	defer func() {
		got = recover()
	}()
	f()
	return
}
```

## Walkthrough

The deferred `panic` overwrites the in-flight panic, masking "original". Dropping it lets the recover capture the real first panic.

## Pitfalls

- A second panic during unwind overwrites the first.
- Keep deferred cleanups panic-free, or recover inside them.
