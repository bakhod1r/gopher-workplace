# Nil map writes panic

## Intuition

A nil map reads as empty but panics on write; construction via `make`/literal is mandatory before assignment.

## Approach

1. A `var m map[string]int` is nil; writing to it panics.
2. Initialize with `make(map[string]int)`.

## Solution

```go
func Tally(words []string) map[string]int {
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}
	return m
}
```

## Walkthrough

The first `m[word]++` on a nil map panics. `make` gives an empty, writable map, so tallying works and `Tally(nil)` returns an empty map.

## Pitfalls

- Reading `m[k]` on a nil map is safe; writing is not.
- `make(map[K]V)` or `map[K]V{}` both work.
