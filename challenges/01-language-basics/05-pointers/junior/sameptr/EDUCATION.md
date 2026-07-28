# Pointer equality is identity

## The idea

`==` on pointers compares the addresses they hold; it is true only when both refer to the same storage.

## Why it matters

Identity checks matter for caches, deduplication, and cycle detection.

## Watch out

- Equal values do NOT imply equal pointers.
- `*a == *b` compares values; `a == b` compares addresses.
