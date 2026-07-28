# Pointers into slice elements

## The idea

`&xs[i]` addresses the actual backing-array slot, so a `[]*int` of these aliases the slice; writes through them mutate xs.

## Why it matters

Editable views and in-place element handles use element pointers.

## Watch out

- Use `&xs[i]` (indexed), not `&v` from a range value in a way that detaches — though Go 1.22 makes `&v` per-iteration too.
- Appending to xs later may reallocate and detach these pointers.
