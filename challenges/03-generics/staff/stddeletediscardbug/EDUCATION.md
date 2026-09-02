# The Deletion That Was Never Kept

## Intuition

`DeleteFunc` compacts the kept elements into the front of the same array and returns a *shorter header*. The array itself still holds the old tail. Using that header only for an emptiness test and then returning the original keeps the original length, so the caller sees the stale tail as live data.

## Approach

1. Return whatever `slices.DeleteFunc` hands back.
2. Accept that the caller's original slice is now scrambled — that is the price of in-place deletion.

## Solution

```go
func Purge[T any](s []T, drop func(T) bool) []T {
	return slices.DeleteFunc(s, drop)
}
```

## Walkthrough

`Purge([1,2,3,4], even)` compacts to `[1 3 3 4]` and returns all four, so `3` appears twice and `4` was never removed at all.

## Pitfalls

- The same mistake with `slices.Compact`, `slices.Insert`, and `append` — all of them return.
- Assuming the discarded elements are zeroed; `DeleteFunc` clears the tail only to release references, and older versions did not.
