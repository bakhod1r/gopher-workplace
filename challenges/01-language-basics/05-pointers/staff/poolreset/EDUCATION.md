# Fully resetting pooled objects

## The idea

Object-pool reuse requires clearing every field; truncating a slice with `[:0]` keeps its capacity for reuse while dropping the length.

## Why it matters

Partial resets leak stale data from a previous use into the next — a real pooling bug.

## Watch out

- Zeroing one field leaves the rest stale.
- `b.Data[:0]` empties the slice but keeps its backing array for reuse.
