# Pointers inside copied struct values

## The idea

Copying a struct duplicates its pointer fields' addresses, so the copy still aliases the same pointee; mutate through the pointer even when the struct itself is a copy.

## Why it matters

Map struct values aren't addressable, but pointer fields inside them still reach shared state.

## Watch out

- `r := m[k]` copies the struct but not the pointee.
- `*r.P++` mutates the shared int.
