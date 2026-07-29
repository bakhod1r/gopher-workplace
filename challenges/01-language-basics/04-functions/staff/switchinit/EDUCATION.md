# The switch init statement

## Intuition

`switch init; tag {}` (or `switch init; {}`) runs the init once and scopes its variables to the switch, avoiding repeated evaluation across cases.

## Approach

1. Calling `compute(x)` in each case re-runs the expensive work.
2. Use a switch init statement: `switch v := compute(x); {` and compare `v`.

## Solution

```go
var calls int

func compute(x int) int { calls++; return x }

func Classify(x int) (label string, runs int) {
	calls = 0
	switch v := compute(x); {
	case v < 0:
		label = "neg"
	case v == 0:
		label = "zero"
	default:
		label = "pos"
	}
	return label, calls
}
```

## Walkthrough

The bug evaluates `compute(x)` once per case. The init form computes it once, binds `v`, and each case tests the cached value.

## Pitfalls

- Cases in a tagless switch each re-evaluate their own expression.
- Hoist shared work into the switch init statement.
