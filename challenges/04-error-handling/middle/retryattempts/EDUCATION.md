# Retry With Attempt Count

## Intuition

A retry loop has three exits: success, exhaustion, and a budget that was never valid. Only the exhaustion path needs to explain itself, and the last failure is the one that explains it.

## Approach

1. Reject `attempts <= 0` before calling `f`.
2. Loop, returning nil on success and remembering the error.
3. Wrap the remembered error with the attempt count.

## Solution

```go
if attempts <= 0 {
	return ErrNoAttempts
}
var last error
for i := 0; i < attempts; i++ {
	last = f()
	if last == nil {
		return nil
	}
}
return fmt.Errorf("after %d attempts: %w", attempts, last)
```

## Walkthrough

When the third call succeeds the loop returns nil immediately, so `f` runs exactly three times and no wrapper is built.

## Pitfalls

- Calling `f` once before the loop, spending an extra attempt.
- Wrapping the first failure instead of the last.
- Running the loop at all when the budget is zero.
