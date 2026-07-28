# Deep-copying pointer fields

## The idea

Struct copy duplicates a pointer field's address, sharing the pointee; a deep clone allocates and copies what the pointer references.

## Why it matters

Shallow clones that share pointer fields cause cross-mutation bugs.

## Watch out

- `*b` copies the pointer value, not the pointee.
- Allocate a new pointee and copy the value into it.
