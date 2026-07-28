# Partitioning into buckets

## The idea

Route each element to one of two accumulators based on a predicate:

```go
for _, x := range xs {
	if x%2 == 0 { evens = append(evens, x) } else { odds = append(odds, x) }
}
```

## Why it matters

Splitting valid/invalid, matched/unmatched, or by category is a universal
pipeline step. One pass, order preserved within each group.

## Watch out

- Named returns start as `nil`; appending is fine, and the test accepts nil/empty.
- Generalizes to N buckets with a `map[key][]T`.
- Order is preserved only within each bucket, not across them.
