# Insert At The End Rejected

## Intuition

Index `len(s)` means "after the last element", which is a legal insertion point; rejecting it turns a valid append into a no-op.

## Approach

1. Reject only `i < 0` and `i > len(s)`.
2. Clone and insert at `i`.

## Solution

```go
func InsertAt[T any](s []T, i int, v T) []T {
	if i < 0 || i > len(s) {
		return s
	}
	return slices.Insert(slices.Clone(s), i, v)
}
```

## Walkthrough

`InsertAt([]int{1}, 1, 2)` should append `2`; the buggy guard sees `1 >= 1` and returns the input.

## Pitfalls

- Reusing an indexing bound (`>= len`) where an insertion bound (`> len`) is needed.
- Special-casing the append instead of fixing the guard.
- Letting the bad index reach `slices.Insert`, which panics.
