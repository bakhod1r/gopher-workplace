# Result Or Zero

## Intuition

When a Go function returns an error, its other results carry no promise. Code that reads them anyway silently mixes garbage into real data.

## Approach

1. Check the error first.
2. Return 0 when it is non-nil.
3. Return the value otherwise.

## Solution

```go
if err != nil {
	return 0
}
return v
```

## Walkthrough

`OrZero(42, ErrHost)` discards the 42 — the host failed, so its reading means nothing.

## Pitfalls

- Returning `v` regardless, which is the bug this helper exists to prevent.
- Checking `v == 0` instead of the error.
- Treating a nil error as failure by inverting the condition.
