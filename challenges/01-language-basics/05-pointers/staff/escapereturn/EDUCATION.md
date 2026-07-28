# Aliasing vs escaping copies

## The idea

`&localCopy` escapes to the heap but is independent of the slice; `&xs[i]` addresses the backing array so writes are visible in the slice.

## Why it matters

Returning a pointer to a copy instead of the element is a subtle no-alias bug.

## Watch out

- `v := xs[i]; &v` gives an independent pointer.
- `&xs[i]` gives an aliasing pointer into the slice.
