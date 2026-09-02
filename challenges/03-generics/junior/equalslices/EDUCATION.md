# Slice Equality

## Intuition

Slices are only comparable to `nil` in Go. Equality has to be defined element by element, and the length check makes the index loop safe.

## Approach

1. Return `false` if the lengths differ.
2. Compare each pair of elements, returning `false` on the first difference.
3. Return `true` after the loop.

## Solution

```go
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
```

## Walkthrough

`Equal([]int{1, 2}, []int{2, 1})` passes the length check, then finds `1 != 2` at index 0 and returns `false`.

## Pitfalls

- Writing `return a == b` — slices are not comparable with `==`.
- Skipping the length check and indexing past the end of the shorter slice.
- Treating `nil` and `[]T{}` as different: both have length 0, so they are equal here.
