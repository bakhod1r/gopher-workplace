# Forwarding variadic arguments

## The idea

`extra...` spreads a slice into another variadic/append call; appending to a caller's slice can clobber it if capacity allows, so copy when purity matters.

## Why it matters

Builder helpers that must not surprise the caller copy before appending.

## Watch out

- `append(base, extra...)` can mutate base's backing array if it has spare cap.
- Copy into a nil-based slice first for a guaranteed-fresh result.
