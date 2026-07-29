# Frequency counting with maps

## Intuition

The map zero value makes counting a one-liner: a missing key reads as 0, so
`m[k]++` reads-modifies-writes without a presence check:

```go
m := make(map[string]int)
for _, w := range words { m[w]++ }
```

## Approach

1. make an empty map[string]int.
2. Range words; result[w]++ leans on the int zero value for the first occurrence.
3. Return the map (empty but non-nil for nil/empty input).

## Solution

```go
func Count(words []string) map[string]int {
	result := make(map[string]int)
	for _, w := range words {
		result[w]++
	}
	return result
}
```

## Walkthrough

Count(["a","b","a"]): "a"->1, "b"->1, "a"->2 -> {"a":2,"b":1}.

## Pitfalls

- `make` the map first; `m[k]++` on a nil map panics (it's a write).
- Iteration order is random — sort keys for deterministic output.
- For non-zero defaults, use comma-ok instead of the zero value.
