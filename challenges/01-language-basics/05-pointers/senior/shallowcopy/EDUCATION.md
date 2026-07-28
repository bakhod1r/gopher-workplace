# Deep-copying tree structures

## The idea

Copying only the root while reusing child pointers aliases the subtrees; a true deep copy recurses on both children.

## Why it matters

Shallow copies that share subtrees corrupt the source on edit.

## Watch out

- `Left: t.Left` shares the original subtree.
- Recurse with `Copy(t.Left)` for independence.
