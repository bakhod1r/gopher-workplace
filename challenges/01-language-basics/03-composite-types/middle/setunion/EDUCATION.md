# Set operations via maps

## Intuition

Insert both inputs into a set (a map), then collect and sort:

```go
set := make(map[int]struct{})
for _, x := range a { set[x] = struct{}{} }
for _, x := range b { set[x] = struct{}{} }
out := make([]int, 0, len(set))
for x := range set { out = append(out, x) }
sort.Ints(out)
```

## Approach

1. Add all of a then all of b into one set.
2. Collect keys into a slice.
3. sort.Ints ascending.
4. Return slice.

## Solution

```go
import "sort"

func Union(a, b []int) []int {
	set := map[int]bool{}
	for _, v := range a {
		set[v] = true
	}
	for _, v := range b {
		set[v] = true
	}
	out := make([]int, 0, len(set))
	for v := range set {
		out = append(out, v)
	}
	sort.Ints(out)
	return out
}
```

## Walkthrough

set from a={3,1,2}, add b -> {1,2,3,4}; sort -> [1,2,3,4].

## Pitfalls

- Map order is random — sort for stable output.
- `struct{}` values use no memory.
- For huge sorted inputs, a merge is more cache-friendly than a map.
