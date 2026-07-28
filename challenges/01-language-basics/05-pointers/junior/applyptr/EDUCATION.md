# In-place transform via a callback

## The idea

Reading `*p`, applying `f`, and storing back mutates the caller's variable with pluggable logic.

## Why it matters

Map-in-place and update hooks combine pointers with function values.

## Watch out

- The pointee is both input and output.
- `f` must be non-nil; calling a nil func panics.
