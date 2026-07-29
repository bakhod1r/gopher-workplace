# Set difference

## Intuition

Build a set from `b`, then keep the elements of `a` absent from it, de-duping:

```go
rm := make(map[int]struct{})
for _, x := range b { rm[x] = struct{}{} }
seen := make(map[int]struct{})
for _, x := range a {
	if _, bad := rm[x]; bad { continue }
	if _, dup := seen[x]; dup { continue }
	seen[x] = struct{}{}; out = append(out, x)
}
sort.Ints(out)
```

## Approach

1. Build set inB from b.
2. For each a-value not in inB, add to a result set (dedupe).
3. Collect keys, sort ascending.
4. Return sorted slice.

## Solution

```go
import "sort"

func Diff(a, b []int) []int {
	inB := map[int]bool{}
	for _, v := range b {
		inB[v] = true
	}
	set := map[int]bool{}
	for _, v := range a {
		if !inB[v] {
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

inB={2,4}. a: 1 keep,2 drop,3 keep,3 dup,4 drop -> set{1,3} -> sort -> [1,3].

## Pitfalls

- De-dup the output; `a` may repeat survivors.
- Sort for stable order.
- The exclusion set makes it O(n+m) instead of nested scans.
