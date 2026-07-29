# tagless switch

## Intuition

A switch with no tag tests each case's boolean condition in order, reading like an if/else-if ladder.

## Approach

1. Use a `switch` with no tag.
2. Each case is a boolean condition, tested in order.

## Solution

```go
func Classify(n int) string {
	switch {
	case n < 0:
		return "negative"
	case n == 0:
		return "zero"
	default:
		return "positive"
	}
}
```

## Walkthrough

`Classify(-5)` matches `n < 0` first → "negative".

## Pitfalls

- Cases are evaluated top to bottom; the first true one runs.
- `default` covers the remaining case.
