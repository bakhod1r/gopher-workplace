# In-place tree transformation

## The idea

Swapping children at every node and recursing mirrors the tree; the swap is a parallel assignment on pointer fields.

## Why it matters

Structural transforms (mirror, prune, rotate) mutate trees recursively.

## Watch out

- Swap before or after recursing — both work since the swap is local.
- Base-case nil returns immediately.
