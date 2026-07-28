# Conditional counting

## The idea

A running counter incremented under an `if` inside a loop is the basis of filters, histograms, and predicates.

## Why it matters

It generalises to any predicate — the shape stays the same as the condition grows.

## Watch out

- Return 0 for no matches, never -1 or a sentinel.
- A nil slice ranges zero times.
