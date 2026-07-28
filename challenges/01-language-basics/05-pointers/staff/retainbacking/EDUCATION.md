# Full-slice expressions and retention

## The idea

`s[:k:k]` limits capacity so appends reallocate instead of spilling; it also lets the unused tail be collected once the parent is gone.

## Why it matters

A small sub-slice pinning a huge array is a real memory leak, and shared spare capacity corrupts data.

## Watch out

- `xs[:k]` keeps cap == cap(xs); appends spill into the parent.
- `xs[:k:k]` (or a copy) isolates the prefix.
