# Indexing needs length, appending needs capacity

## Intuition

`make([]int, 0, n)` reserves capacity but has length 0 — indexing `out[i]` panics.
Two valid patterns:

```go
out := make([]int, len(xs)); for i, x := range xs { out[i] = x * 2 } // index
out := make([]int, 0, len(xs)); for _, x := range xs { out = append(out, x*2) } // append
```

## Approach

1. Bug: make([]int, 0, len(xs)) has length 0, so out[i] = x*2 indexes out of range and panics. 2. Fix: make([]int, len(xs)) so the slice has full length and index writes are valid. 3. Each out[i] then receives the doubled value.

## Solution

```go
func Doubled(xs []int) []int {
	out := make([]int, len(xs))
	for i, x := range xs {
		out[i] = x * 2
	}
	return out
}
```

## Walkthrough

With length 0, out[0]=... panics (index out of range). With length len(xs), out[0..2] are valid slots filled with 2,4,6.

## Pitfalls

- Capacity is spare room; only elements within length exist.
- `append` grows length; indexing requires it to already be there.
- `make([]T, n)` zero-fills n elements; `make([]T, 0, n)` fills none.
