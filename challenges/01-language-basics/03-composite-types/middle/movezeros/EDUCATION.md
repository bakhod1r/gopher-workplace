# Stable compaction

## Intuition

Collect the non-zero elements in order, then pad the rest with zeros:

```go
out := []int{}
for _, x := range xs { if x != 0 { out = append(out, x) } }
for len(out) < len(xs) { out = append(out, 0) }
```

## Approach

1. Walk xs; count zeros, append non-zeros in order.
2. After the pass, append that many zeros.
3. Return the new slice.

## Solution

```go
func MoveZeros(xs []int) []int {
	out := make([]int, 0, len(xs))
	zeros := 0
	for _, v := range xs {
		if v == 0 {
			zeros++
		} else {
			out = append(out, v)
		}
	}
	for ; zeros > 0; zeros-- {
		out = append(out, 0)
	}
	return out
}
```

## Walkthrough

[0,1,0,3,12]: non-zeros [1,3,12], zeros=2 -> append 0,0 -> [1,3,12,0,0].

## Pitfalls

- Preserve the **order** of non-zeros (stable) — a swap-based version wouldn't.
- Output length equals input length.
- An in-place two-pointer version avoids the second allocation.
