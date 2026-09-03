# Emptying A Map Without Rebuilding It

## Intuition

A map value is a pointer to a header. Assigning a new map replaces your pointer; clearing the map changes what everyone's pointer points at.

## Approach

1. `Reset` calls `clear(m)`, which is defined as a no-op for a nil map.
2. `Tally` resets, then counts.

## Solution

```go
func Reset(m map[string]int) {
	clear(m)
}

func Tally(m map[string]int, words []string) map[string]int {
	Reset(m)
	for _, w := range words {
		m[w]++
	}
	return m
}
```

## Walkthrough

Because `Tally` writes into the map the caller passed in, a map sized once at startup serves every later round without touching the allocator.

## Pitfalls

- `m = make(map[string]int)` inside `Reset`, which the caller never sees.
- Deleting keys while ranging when a plain `clear` would do.
- Reusing a map that briefly held a huge number of keys — the buckets stay allocated for the lifetime of the map.
