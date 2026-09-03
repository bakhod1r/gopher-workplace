# Retry Only What Is Retryable

## Intuition

Retrying is only correct for failures that might not recur. Retrying a rejected request wastes the budget and multiplies load for a request that can never succeed.

## Approach

1. Loop up to `attempts` times.
2. Return nil on success.
3. Return the error immediately when it is not transient; otherwise continue.
4. Return the last error after the budget is spent.

## Solution

```go
var last error
for i := 0; i < attempts; i++ {
	last = f()
	if last == nil {
		return nil
	}
	if !errors.Is(last, ErrTransient) {
		return last
	}
}
return last
```

## Walkthrough

The transient case wraps the sentinel with an attempt number and still matches, so the loop keeps going until the third call succeeds.

## Pitfalls

- Retrying every failure, including permanent ones.
- Using `==` so a wrapped transient error is treated as fatal.
- Returning nil when the budget is exhausted.
