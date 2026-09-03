# Sizing The Map Up Front

## Intuition

An unhinted map starts with room for a handful of keys. Every time it fills, the runtime allocates a bigger bucket array and rehashes what is already there.

## Approach

1. `make(map[string]int, len(words))`.
2. Range with the index; write only when the key is missing.

## Solution

```go
func Index(words []string) map[string]int {
	m := make(map[string]int, len(words))
	for i, w := range words {
		if _, ok := m[w]; !ok {
			m[w] = i
		}
	}
	return m
}
```

## Walkthrough

For 1000 distinct keys, the hintless version allocates around 13 times while the hinted one lands near 6 — the difference is the growth staircase, and it gets worse as the map gets bigger.

## Pitfalls

- `m[w] = i` unconditionally, which records the *last* occurrence.
- Passing an over-large hint, which allocates buckets you never fill.
- Assuming the hint prevents growth; it only front-loads it.
