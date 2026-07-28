# Struct identity via pointers

## The idea

Pointer `==` compares addresses, distinguishing two structurally-equal instances.

## Why it matters

Identity checks power caches, sets of objects, and cycle detection.

## Watch out

- `*a == *b` compares fields; `a == b` compares identity.
- Distinct `&Cart{}` values are never equal pointers.
