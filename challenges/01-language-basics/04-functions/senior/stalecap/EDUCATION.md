# Element pointers and reallocation

## The idea

`append` past capacity moves the backing array; any pointer/slice into the old array is now detached from the live slice.

## Why it matters

Holding element pointers across appends is a real aliasing bug in buffers and pools.

## Watch out

- After a reallocating append, old `&xs[i]` is stale.
- Index the current slice (`xs[0]`) or re-take the address.
