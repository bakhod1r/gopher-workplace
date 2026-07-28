# Rank by frequency

## The idea

Count with a map, then sort the distinct keys by a compound key — count
descending, then the value ascending for stable tie-breaks:

```go
sort.Slice(keys, func(i, j int) bool {
	if m[keys[i]] != m[keys[j]] { return m[keys[i]] > m[keys[j]] }
	return keys[i] < keys[j]
})
```

## Why it matters

Leaderboards, "most common", and heavy-hitter detection are count-then-rank.
Deterministic tie-breaking keeps output stable for tests and users.

## Watch out

- Map iteration is random; you must sort for a stable ranking.
- Break ties explicitly, or equal counts come out in arbitrary order.
- For large data, a heap of size k beats a full sort.
