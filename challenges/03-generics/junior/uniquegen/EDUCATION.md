# Unique

## Intuition

A linear scan with a seen-set is O(n). The map key type is the type parameter itself, which the `comparable` constraint makes legal.

## Approach

1. Create a `seen` map and an `out` slice.
2. For each element, skip it when already seen.
3. Otherwise mark it seen and append it.

## Solution

```go
func Unique[T comparable](s []T) []T {
	seen := make(map[T]bool, len(s))
	out := make([]T, 0, len(s))
	for _, e := range s {
		if seen[e] {
			continue
		}
		seen[e] = true
		out = append(out, e)
	}
	return out
}
```

## Walkthrough

`Unique([]int{1, 2, 1})` appends 1, appends 2, then finds `seen[1]` already true and skips the third element.

## Pitfalls

- Using `[T any]` — a type parameter must be `comparable` to be a map key.
- Sorting first, which destroys first-seen order.
- Checking `out` with a nested loop (O(n²)) when a map is available.
