# Recursion over binary trees

## The idea

Tree algorithms recurse on Left and Right, base-casing on nil; height combines the deeper subtree.

## Why it matters

Depth, balance, and traversal all follow this recursive shape.

## Watch out

- Base-case nil before dereferencing.
- Height uses max of subtrees; size would use sum.
