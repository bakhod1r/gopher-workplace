# Merging with a dummy head

## The idea

A sentinel dummy node removes special-casing the first append; a tail pointer builds the merged list in order.

## Why it matters

Merge is the core of merge sort and ordered stream joins.

## Watch out

- The dummy head avoids branching on the first node.
- Attach the leftover list directly once one runs out.
