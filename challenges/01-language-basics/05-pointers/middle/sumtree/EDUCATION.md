# Aggregating over a tree

## The idea

A post-order recursion combines the current value with both subtree sums.

## Why it matters

Weighted totals and reductions over trees follow this pattern.

## Watch out

- Base-case nil returns the identity (0 for sums).
- Recurse both children.
