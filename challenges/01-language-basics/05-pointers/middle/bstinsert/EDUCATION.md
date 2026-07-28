# Building a binary search tree

## The idea

Insertion recurses to the correct nil slot and returns each subtree so parents can relink; the root return covers the empty tree.

## Why it matters

BSTs underpin ordered sets, maps, and range queries.

## Watch out

- Returning the subtree lets the parent reattach it.
- The nil slot is where the new node goes.
