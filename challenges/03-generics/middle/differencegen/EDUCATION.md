# Difference

## Intuition

Two maps look redundant until you notice they answer different questions: "is it in `b`?" and "have I emitted it already?".

## Approach

1. Index `b` into an exclusion set.
2. Walk `a`, skipping excluded or already-emitted values.
3. Append the rest.

## Solution

```go
func Difference[T comparable](a, b []T) []T {
	exclude := make(map[T]bool, len(b))
	for _, v := range b {
		exclude[v] = true
	}
	seen := make(map[T]bool, len(a))
	out := make([]T, 0, len(a))
	for _, v := range a {
		if exclude[v] || seen[v] {
			continue
		}
		seen[v] = true
		out = append(out, v)
	}
	return out
}
```

## Walkthrough

`Difference([]int{1,2,2,3}, []int{2})` drops both `2`s via the exclusion set and emits `1` and `3`.

## Pitfalls

- Scanning `b` inside the loop over `a`, making the function quadratic.
- Emitting duplicates from `a`.
- Sorting the output and losing first-seen order.
