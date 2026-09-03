# Open Circuit

## Intuition

A breaker converts a slow, repeated failure into a fast, cheap one. The decision must happen before the call, otherwise the dependency is still being hit.

## Approach

1. Return `ErrOpen` when `failures >= Threshold`.
2. Call `f`, resetting the counter on success.
3. Increment and return the error on failure.

## Solution

```go
if b.failures >= b.Threshold {
	return ErrOpen
}
if err := f(); err != nil {
	b.failures++
	return err
}
b.failures = 0
return nil
```

## Walkthrough

With a threshold of 0 the guard fires on the very first call, so `f` never runs — the circuit starts open.

## Pitfalls

- Calling `f` before checking the state, so the dependency still gets traffic.
- Incrementing instead of resetting on success, so the circuit eventually opens regardless.
- Wrapping the failure in `ErrOpen`, making a real failure look like a rejection.
