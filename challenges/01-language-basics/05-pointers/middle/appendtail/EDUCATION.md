# Updating a head pointer via double indirection

## The idea

Appending to an empty list must reseat the caller's head; a `**Node` lets the callee do that, while non-empty lists just link at the tail.

## Why it matters

Generic list-append helpers take a pointer-to-head to cover the empty case.

## Watch out

- Without `**Node`, an append to an empty list can't update the caller.
- Walk `n.Next` until nil to find the tail.
