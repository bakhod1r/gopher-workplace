# Map as an index

## The idea

Walk once; for each element look up whether its complement was already seen. A
map from value to index gives O(1) lookups:

```go
seen := make(map[int]int)
for i, x := range xs {
	if j, ok := seen[target-x]; ok { return j, i, true }
	seen[x] = i
}
return 0, 0, false
```

## Why it matters

Turning an O(n²) nested scan into O(n) with a lookup map is a fundamental
technique — the same idea behind joins and dedup.

## Watch out

- Check for the complement **before** inserting the current element, or you may
  pair an element with itself.
- The map stores the earliest index of each value.
- Returns the first pair found; multiple pairs may exist.
