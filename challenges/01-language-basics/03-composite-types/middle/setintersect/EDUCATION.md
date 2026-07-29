# Intersection with a lookup set

## Intuition

Put one side in a set, then filter the other by membership, de-duplicating the
result:

```go
in := make(map[int]struct{})
for _, x := range a { in[x] = struct{}{} }
seen := make(map[int]struct{})
for _, x := range b {
	if _, ok := in[x]; ok {
		if _, dup := seen[x]; !dup { out = append(out, x); seen[x] = struct{}{} }
	}
}
sort.Ints(out)
```

## Approach

1. Build set inA from a.
2. For each b-value in inA, add to result set (dedupe).
3. Collect keys, sort ascending.
4. Return slice.

## Solution

```go
import "sort"

func Intersect(a, b []int) []int {
	inA := map[int]bool{}
	for _, v := range a {
		inA[v] = true
	}
	set := map[int]bool{}
	for _, v := range b {
		if inA[v] {
			set[v] = true
		}
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

inA={1,2,3,4}. b: 2 in ->set{2}; 4 in ->{2,4}; 6 no; 2 dup. sort -> [2,4].

## Pitfalls

- De-dup the output: `b` may repeat a common value.
- Membership is `_, ok := set[x]`.
- Sort for deterministic order.
