# Bucket Counts

## Intuition

Using only `<` matters: it means the same code works for any ordered type, and the half-open convention guarantees each element is counted exactly once.

## Approach

1. Allocate `len(edges)+1` counters.
2. For each element, advance past every edge it is not below.
3. Increment that bucket.

## Solution

```go
func Buckets[T cmp.Ordered](s []T, edges []T) []int {
	out := make([]int, len(edges)+1)
	for _, v := range s {
		i := 0
		for i < len(edges) && !(v < edges[i]) {
			i++
		}
		out[i]++
	}
	return out
}
```

## Walkthrough

`Buckets([]int{1,5,9}, []int{5})` puts `1` below the edge and both `5` and `9` at or above it.

## Pitfalls

- Allocating `len(edges)` counters and dropping the final bucket.
- Using `<=` on the edge, which puts boundary values in the wrong bucket.
- Assuming the edges are sorted without saying so in the doc.
