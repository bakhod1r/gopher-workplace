# Slice index bounds

## The idea

Valid indices run `0 .. len(xs)-1`; seeding a pointer at `len(xs)` and dereferencing it panics immediately.

## Why it matters

Off-by-one initial bounds are a classic index-out-of-range crash in two-pointer code.

## Watch out

- The highest valid index is `len(xs)-1`, not `len(xs)`.
- `xs[len(xs)]` always panics.
