# Argmax over a count map

## The idea

Count occurrences, then select the key with the highest count, breaking ties
deterministically (here, the smaller value):

```go
best, bestN, have := 0, 0, false
for v, n := range counts {
	if !have || n > bestN || (n == bestN && v < best) { best, bestN, have = v, n, true }
}
```

## Why it matters

Mode, plurality winner, and heavy-hitter selection are argmax-over-counts. Because
map order is random, an explicit tie-break is required for reproducibility.

## Watch out

- Map iteration order is random — never rely on it; encode the tie-break.
- Return `ok=false` for empty rather than a fake value.
- Initialize the "best" carefully (use a `have` flag).
