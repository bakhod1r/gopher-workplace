# Aliasing vs copying slice views

## The idea

Multiple slices of the same array share storage; copying into a new slice breaks that aliasing.

## Why it matters

Whether views alias determines if a write in one is visible in another — a key correctness property.

## Watch out

- `p[:]` aliases the array; `append([]int(nil), p[:]...)` copies it.
- Aliased views see each other's writes.
