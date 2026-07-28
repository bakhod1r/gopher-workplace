# Initialising nested maps

## The idea

Only the outer map exists after `map[K]map[..]{}`; each inner map is nil until explicitly created, and writing to a nil map panics.

## Why it matters

Adjacency lists, indexes, and grouped counters all hit this nil-inner-map crash.

## Watch out

- `g[from]` is nil, not an empty map, until you assign one.
- Lazily create the inner map before writing to it.
