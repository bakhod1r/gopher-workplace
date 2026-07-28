# Stable compaction

## The idea

Collect the non-zero elements in order, then pad the rest with zeros:

```go
out := []int{}
for _, x := range xs { if x != 0 { out = append(out, x) } }
for len(out) < len(xs) { out = append(out, 0) }
```

## Why it matters

Partition/compact-while-preserving-order shows up in memory management, sparse
data, and UI reordering. The two-phase build is the simplest stable version.

## Watch out

- Preserve the **order** of non-zeros (stable) — a swap-based version wouldn't.
- Output length equals input length.
- An in-place two-pointer version avoids the second allocation.
