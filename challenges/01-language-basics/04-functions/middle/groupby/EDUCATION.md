# Bucketing with a key function

## Intuition

`append(m[k], x)` relies on a missing key yielding a nil slice that append grows — no pre-init per key needed.

## Approach

1. Init a `map[int][]int`.
2. Append each element to the bucket for `key(x)`.

## Solution

```go
func GroupBy(xs []int, key func(int) int) map[int][]int {
	m := map[int][]int{}
	for _, x := range xs {
		k := key(x)
		m[k] = append(m[k], x)
	}
	return m
}
```

## Walkthrough

With `mod2`, evens land in bucket 0 and odds in bucket 1.

## Pitfalls

- You must init the outer map; per-key slices need no init.
- Input order is preserved within each bucket by appending in order.
