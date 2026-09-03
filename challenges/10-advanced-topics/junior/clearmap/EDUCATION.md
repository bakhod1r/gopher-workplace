# Empty The Map, Keep The Map

## Intuition

The map parameter is a copy of the caller's map *handle*. Writing through the handle (deleting keys) is visible to the caller; overwriting the handle is not.

## Approach

1. Call the builtin `clear` on the map.
2. `clear` on a nil map is a no-op, so nothing extra is needed.

## Solution

```go
// Reset removes every entry from m without replacing the map.
//
// Callers hold the same map value, so the entries must be deleted in place
// rather than by assigning a fresh map.
//
// Examples:
//
// 	m := map[string]int{"a": 1}; Reset(m) => len(m) == 0
func Reset(m map[string]int) {
	clear(m)
}
```

## Walkthrough

`m := map[string]int{"a":1}; alias := m` — two handles, one table. `clear(m)` empties the table, so `len(alias) == 0`. `m = map[string]int{}` would give `m` a new table and leave `alias` holding the old one.

## Pitfalls

- Assigning a new map inside the function — the caller never sees it.
- Ranging and deleting works, but `clear` says it in one line and handles nil.
