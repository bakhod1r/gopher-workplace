# Deterministic order from a map

## Intuition

Go randomizes map iteration order on purpose. To get a stable order, collect the
keys into a slice and sort:

```go
out := make([]string, 0, len(m))
for k := range m { out = append(out, k) }
sort.Strings(out)
```

## Approach

1. Allocate a slice with capacity len(m).
2. Range the map, appending each key (map order is random).
3. sort.Strings on the collected keys.
4. Return the sorted slice.

## Solution

```go
import "sort"

func Sorted(m map[string]int) []string {
	result := make([]string, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}
```

## Walkthrough

Sorted({"banana","apple","cherry"}): collect in random order, sort.Strings -> ["apple","banana","cherry"].

## Pitfalls

- Pre-size with `make(..., 0, len(m))` to avoid regrowth.
- Ranging a map with only `k` iterates keys.
- `slices.Sorted(maps.Keys(m))` (Go 1.23+) does this in one line.
