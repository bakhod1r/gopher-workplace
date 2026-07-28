# Traversing pointer-linked structures

## The idea

Following `Next` pointers to nil visits every node; nil marks the end.

## Why it matters

Linked lists, trees, and chains are all pointer traversals.

## Watch out

- Stop at nil, don't dereference it.
- An empty list (nil head) has length 0.
