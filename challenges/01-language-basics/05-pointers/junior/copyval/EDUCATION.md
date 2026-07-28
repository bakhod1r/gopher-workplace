# Copying values via pointers

## The idea

`*dst = *src` reads the source pointee and writes it to the destination pointee — a value copy, not an alias.

## Why it matters

Assigning between referenced variables is a building block of in-place algorithms.

## Watch out

- After the copy the two remain separate ints.
- Order matters: dst is the target.
