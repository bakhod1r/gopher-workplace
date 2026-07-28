# Searching a BST

## The idea

The ordering invariant lets search discard one subtree per step, giving O(height) lookups.

## Why it matters

Ordered lookups and range membership use BST search.

## Watch out

- Base-case nil returns false.
- Compare then recurse into exactly one child.
