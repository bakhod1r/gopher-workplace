# Tagless switch as an if-ladder

## The idea

`switch {}` with boolean case expressions is a clean alternative to long if/else-if chains; Go breaks after the first match automatically.

## Why it matters

It flattens multi-branch logic and avoids the C-style fallthrough footgun by default.

## Watch out

- Cases evaluate top-down; put narrower conditions first if they overlap.
- Go does NOT fall through unless you write `fallthrough`.
