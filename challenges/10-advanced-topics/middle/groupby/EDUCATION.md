# Bucket Pairs Without Growing Every Bucket

## Intuition

Growth is only expensive when the size is a surprise. One cheap counting pass turns thousands of doubling chains into one exactly-sized allocation each.

## Approach

1. Count occurrences per key.
2. Allocate the result map sized to the distinct-key count, and each bucket sized to its count.
3. Walk the pairs again, appending into the pre-sized buckets.

## Solution

```go
// Group collects the second element of each pair into a bucket keyed by
// the first, preserving input order within a bucket.
//
// Both the map and each bucket should be sized from what the input already
// tells you, instead of growing from nothing.
//
// Examples:
//
// 	Group([][2]int{{1, 10}, {1, 11}, {2, 20}}) => map[1:[10 11] 2:[20]]
func Group(pairs [][2]int) map[int][]int {
	if len(pairs) == 0 {
		return map[int][]int{}
	}
	counts := make(map[int]int, len(pairs))
	for _, p := range pairs {
		counts[p[0]]++
	}
	out := make(map[int][]int, len(counts))
	for k, n := range counts {
		out[k] = make([]int, 0, n)
	}
	for _, p := range pairs {
		out[p[0]] = append(out[p[0]], p[1])
	}
	return out
}
```

## Walkthrough

With 300 pairs over 5 keys, the counting pass finds 60 per key; each bucket is allocated once at capacity 60 and never grows.

## Pitfalls

- Returning nil for an empty input, which is a different value from an empty map.
- Sizing the buckets to `len(pairs)`, which is correct and wastes most of it.
