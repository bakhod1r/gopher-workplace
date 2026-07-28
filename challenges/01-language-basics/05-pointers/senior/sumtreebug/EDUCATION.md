# Recursing into all children

## The idea

A tree aggregation must combine the node with BOTH subtree results; returning only the node value discards the rest.

## Why it matters

Forgetting the recursive calls silently returns a partial aggregate.

## Watch out

- `return t.Val` visits one node only.
- Add `SumTree(t.Left) + SumTree(t.Right)`.
