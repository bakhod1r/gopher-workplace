# Full-slice expressions and capacity

## The idea

`s[low:high:max]` sets capacity to `max-low`; using it to cap a sub-slice prevents appends from reaching the parent and lets the unused tail be garbage-collected.

## Why it matters

Retaining a big array via a small sub-slice is a real memory leak, and shared spare capacity is a real corruption source.

## Watch out

- `xs[:k]` keeps cap == cap(xs); appends spill into the parent.
- `xs[:k:k]` (or a copy) isolates the head.
