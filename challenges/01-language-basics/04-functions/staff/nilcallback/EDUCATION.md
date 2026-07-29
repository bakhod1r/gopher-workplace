# Nil function values

## Intuition

A func variable defaults to nil and panics on call; optional callbacks must be checked before invocation.

## Approach

1. Calling a nil `hook` panics.
2. Guard: use `hook(v)` when non-nil, else add `v` unchanged.

## Solution

```go
func Process(xs []int, hook func(int) int) int {
	total := 0
	for _, v := range xs {
		if hook != nil {
			total += hook(v)
		} else {
			total += v
		}
	}
	return total
}
```

## Walkthrough

`Process(xs, nil)` panics calling the nil function. The guard sums raw values when no hook is provided, giving 6.

## Pitfalls

- A nil func value is not callable — check `hook != nil` first.
- Provide a sensible default (identity) for the nil case.
