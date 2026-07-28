# Shallow vs deep struct copies

## The idea

Copying a struct duplicates value fields but shares slices/maps/pointers; a deep clone must copy those reference fields explicitly.

## Why it matters

Shallow clones that share slices/maps are a frequent aliasing bug.

## Watch out

- `*d` copies the header, not the slice's elements.
- Deep-copy each reference-typed field.
