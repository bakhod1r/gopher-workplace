# Counting with the map zero value

## Intuition

Reading a missing map key yields the value type's zero — 0 for ints. So
incrementing works without checking presence:

```go
m := make(map[string]int)
for _, x := range xs { m[x]++ }
```

`m[x]++` reads (0 if absent), adds one, and stores.

## Approach

1. make an empty map[string]int.
2. Range over xs; result[x]++ relies on the int zero value 0 for first sightings.
3. Return the map.

## Solution

```go
func Count(xs []string) map[string]int {
	result := make(map[string]int)
	for _, x := range xs {
		result[x]++
	}
	return result
}
```

## Walkthrough

Count(["a","b","a"]): "a"->1, "b"->1, "a"->2. Result {"a":2,"b":1}.

## Pitfalls

- You must `make` the map first; `m[x]++` on a nil map panics (it's a write).
- Only counts are convenient; for non-zero defaults you still need comma-ok.
- Iteration order of the result is random.
