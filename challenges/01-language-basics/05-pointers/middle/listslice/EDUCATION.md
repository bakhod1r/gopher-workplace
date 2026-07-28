# Materialising a list

## The idea

Traversing and appending converts a pointer chain to a contiguous slice.

## Why it matters

Serialising or indexing a list often starts by materialising it.

## Watch out

- Return empty (len 0) for nil head.
- One pass, append each value.
