# The for-range loop

## The idea

`range` over a slice returns `(index, value)` copies; discarding the index with `_` gives a clean value iteration.

## Why it matters

It's the default way to traverse slices, maps, strings, and channels in Go.

## Watch out

- `v` is a copy; writing to it does not change the slice.
- A nil slice ranges zero times — no special case needed.
