# Nil map writes panic

## The idea

A nil map reads as empty but panics on write; construction via `make`/literal is mandatory before assignment.

## Why it matters

Returning an un-made map from a constructor is a frequent nil-map-write crash.

## Watch out

- Reading `m[k]` on a nil map is safe; writing is not.
- `make(map[K]V)` or `map[K]V{}` both work.
