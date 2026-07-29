# nil slices are valid append targets

## Intuition

A nil slice has len 0 and appends normally; special-casing it is unnecessary and here actively wrong.

## Approach

1. `append` to a nil slice is valid and returns a fresh slice.
2. Remove the early `if xs == nil { return nil }` that drops the extras.

## Solution

```go
func Collect(xs []int, extra []int) []int {
	for _, v := range extra {
		xs = append(xs, v)
	}
	return xs
}
```

## Walkthrough

The guard returns nil before appending, losing the `extra` elements. Deleting it lets `append` build the combined slice from a nil base.

## Pitfalls

- `append(nil, ...)` allocates a fresh slice — no pre-init needed.
- Ranging or `len` on a nil slice is also fine.
