# Contains

## Intuition

A type parameter only supports the operations its constraint promises. `any` promises nothing, so `e == v` is a compile error; `comparable` promises `==` and `!=`.

## Approach

1. Range over `s`.
2. Return `true` on the first element equal to `v`.
3. Return `false` after the loop.

## Solution

```go
func Contains[T comparable](s []T, v T) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}
```

## Walkthrough

`Contains([]int{1, 2, 3}, 2)` compares `1 == 2` (no), then `2 == 2` (yes) and returns immediately without touching the third element.

## Pitfalls

- Using `[T any]` — `invalid operation: e == v (incomparable types in type set)`.
- Returning `false` inside the loop, which stops after the first element.
- Comparing indexes instead of values (`i == v`).
