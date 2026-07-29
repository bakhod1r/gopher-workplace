# Tagless switch as an if-ladder

## Intuition

`switch {}` with boolean case expressions is a clean alternative to long if/else-if chains; Go breaks after the first match automatically.

## Approach

1. A tagless `switch` tests boolean cases in order.
2. Negative, zero, else positive.

## Solution

```go
func Sign(n int) string {
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

`Sign(-4)` matches `n < 0` → "negative".

## Pitfalls

- Cases evaluate top-down; put narrower conditions first if they overlap.
- Go does NOT fall through unless you write `fallthrough`.
