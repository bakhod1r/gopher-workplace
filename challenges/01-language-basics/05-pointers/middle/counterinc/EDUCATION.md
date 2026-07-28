# Pointer receivers mutate

## The idea

A pointer-receiver method operates on the caller's value; a value-receiver method gets a copy and can't persist mutations.

## Why it matters

Any method that changes struct state needs a pointer receiver.

## Watch out

- `(c Counter)` (value) would increment a copy.
- Go auto-takes the address for `c.Inc()` when `c` is addressable.
