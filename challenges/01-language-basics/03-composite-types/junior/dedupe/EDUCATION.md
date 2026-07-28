# Deduplication

## The idea

Remove duplicates by tracking what you've seen. A set (`map[T]struct{}`) gives
O(1) membership and preserves first-seen order when you append on first sight:

```go
if _, ok := seen[x]; !ok { out = append(out, x); seen[x] = struct{}{} }
```

## Why it matters

Deduping lists (tags, IDs, events) is everywhere. The set-based approach is O(n)
versus O(n²) for a nested scan.

## Watch out

- Emit on **absence**, then record.
- Order-preserving dedup differs from sort-then-compact.
- Element type must be comparable to be a set key.
