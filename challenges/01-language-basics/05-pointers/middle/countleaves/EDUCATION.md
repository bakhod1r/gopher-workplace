# Classifying nodes during traversal

## The idea

A leaf test (both children nil) plus a recursive sum counts leaves; the same shape counts internal nodes or full nodes.

## Why it matters

Structural tree metrics recurse and aggregate over children.

## Watch out

- Distinguish nil (0) from a leaf (1).
- Sum both subtrees for internal nodes.
