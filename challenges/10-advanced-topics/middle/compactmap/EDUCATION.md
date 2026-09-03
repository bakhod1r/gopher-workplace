# A Map That Gives Its Buckets Back

## Intuition

A Go map's bucket array only ever grows. Deleting entries leaves the buckets allocated and empty, so the only way to hand the memory back is to build a new map at the size you now need.

## Approach

1. `make(map[string]int, len(m))`.
2. Range over `m` and copy every entry.
3. Return the new map.

## Solution

```go
// Compact returns a new map holding the same entries as m, sized to the
// entries it actually keeps.
//
// A map that grew to millions of entries keeps its bucket array after the
// entries are deleted; rebuilding is the only way to release it.
//
// Examples:
//
// 	Compact(map[string]int{"a": 1}) => a new map[a:1]
func Compact(m map[string]int) map[string]int {
	out := make(map[string]int, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}
```

## Walkthrough

A map grown to 4096 entries and cut to 512 still holds the 4096-bucket array. Copying the 512 survivors into a map sized 512 lets the big array be collected.

## Pitfalls

- Returning `m` itself when it is already small — the caller asked for a rebuild, and identity is what the test checks.
- Forgetting the size hint and rehashing all the way back up.
