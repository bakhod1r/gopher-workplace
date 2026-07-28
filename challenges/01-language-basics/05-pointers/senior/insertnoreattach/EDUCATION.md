# Reattaching returned subtrees

## The idea

Recursive tree mutators return the (possibly new) subtree; the parent must reassign it, or newly created nodes are lost.

## Why it matters

Forgetting to reattach is a silent data-loss bug in immutable-style tree code.

## Watch out

- `Insert(root.Left, v)` alone drops the result for empty subtrees.
- Always write `root.Left = Insert(root.Left, v)`.
