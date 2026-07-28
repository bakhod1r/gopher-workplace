# Append and shared backing arrays

## The idea

Appending to a slice with spare capacity mutates shared memory; a full-slice expression `s[:len:len]` sets cap=len so the next append copies.

## Why it matters

This aliasing bug corrupts data when one buffer is the base for several derived slices.

## Watch out

- `append` reuses capacity when available — two appends to the same base collide.
- `base[:n:n]` forces the next append to allocate fresh.
