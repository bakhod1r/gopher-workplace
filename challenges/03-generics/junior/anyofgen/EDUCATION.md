# Any and All

## Intuition

`Any` looks for one witness; `All` looks for one counterexample. Both stop at the first decisive element, which is why the empty-slice answers differ.

## Approach

1. `Any`: return `true` at the first accepted element, `false` after the loop.
2. `All`: return `false` at the first rejected element, `true` after the loop.

## Solution

```go
func Any[T any](s []T, pred func(T) bool) bool {
	for _, e := range s {
		if pred(e) {
			return true
		}
	}
	return false
}

func All[T any](s []T, pred func(T) bool) bool {
	for _, e := range s {
		if !pred(e) {
			return false
		}
	}
	return true
}
```

## Walkthrough

`All([]int{}, isEven)` never enters the loop and falls through to `return true` — no element violates the predicate.

## Pitfalls

- Making `All` return `false` for an empty slice.
- Making `Any` return `true` for an empty slice.
- Accumulating a bool across the whole slice instead of returning early — correct, but it loses short-circuiting.
