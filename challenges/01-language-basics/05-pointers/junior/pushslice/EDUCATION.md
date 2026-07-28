# Growing a caller's slice

## The idea

Because append returns a possibly-new header, mutating the caller's slice requires assigning the result through a `*[]int`.

## Why it matters

Builder helpers that must grow the caller's slice take a slice pointer.

## Watch out

- A plain `[]int` parameter can't propagate a reallocated header.
- `*sp = append(*sp, v)` updates the caller.
