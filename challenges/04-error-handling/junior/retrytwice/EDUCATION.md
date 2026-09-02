# Retry Once

## Intuition

A retry is a loop with a hard bound. Because the function is an ordinary value, the retry logic knows nothing about what it retries.

## Approach

1. Call `f` once and return nil on success.
2. Call `f` again.
3. Return whatever the second call returned.

## Solution

```go
if err := f(); err == nil {
	return nil
}
return f()
```

## Walkthrough

When the first call fails and the second succeeds, `f` ran twice and the returned error is nil.

## Pitfalls

- Calling `f` a third time — two attempts total is the contract.
- Returning the *first* error, hiding what actually happened last.
- Calling `f` twice unconditionally, doubling the work on success.
