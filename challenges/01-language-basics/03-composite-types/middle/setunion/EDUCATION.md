# Set operations via maps

## The idea

Insert both inputs into a set (a map), then collect and sort:

```go
set := make(map[int]struct{})
for _, x := range a { set[x] = struct{}{} }
for _, x := range b { set[x] = struct{}{} }
out := make([]int, 0, len(set))
for x := range set { out = append(out, x) }
sort.Ints(out)
```

## Why it matters

Union/intersection/difference on tags, permissions, or IDs are everyday set
tasks. Maps give O(1) membership; sorting restores a deterministic order.

## Watch out

- Map order is random — sort for stable output.
- `struct{}` values use no memory.
- For huge sorted inputs, a merge is more cache-friendly than a map.
