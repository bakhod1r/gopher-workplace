# Handling leftovers in a merge

## The idea

The merge loop stops when either list empties; the non-empty remainder must be linked to the tail to avoid dropping nodes.

## Why it matters

Forgetting the remainder truncates merged output — a real merge-sort bug.

## Watch out

- The loop condition `a != nil && b != nil` leaves one list non-empty.
- Link the leftover directly; it's already sorted.
