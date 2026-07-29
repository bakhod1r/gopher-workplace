# Initialising nested maps

## Intuition

Only the outer map exists after `map[K]map[..]{}`; each inner map is nil until explicitly created, and writing to a nil map panics.

## Approach

1. Writing `g[from][to]` when the inner map is nil panics.
2. Lazily create it: `if g[from] == nil { g[from] = map[string]bool{} }`.

## Solution

```go
func Add(g map[string]map[string]bool, from, to string) {
	if g[from] == nil {
		g[from] = map[string]bool{}
	}
	g[from][to] = true
}
```

## Walkthrough

The first edge from a new node finds a nil inner map and panics. Initializing it first makes the assignment safe.

## Pitfalls

- `g[from]` is nil, not an empty map, until you assign one.
- Lazily create the inner map before writing to it.
