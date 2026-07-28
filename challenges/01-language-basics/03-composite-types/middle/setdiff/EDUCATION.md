# Set difference

## The idea

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

## Why it matters

Difference completes the set-ops trio (union/intersect/diff) used for
permissions, diffs, and reconciliation.

## Watch out

- De-dup the output; `a` may repeat survivors.
- Sort for stable order.
- The exclusion set makes it O(n+m) instead of nested scans.
