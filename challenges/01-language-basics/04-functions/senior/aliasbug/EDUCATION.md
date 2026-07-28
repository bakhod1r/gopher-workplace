# Slice assignment copies only the header

## The idea

Assigning a slice duplicates its (ptr,len,cap) header but not the backing array, so writes through either alias are shared.

## Why it matters

This aliasing silently mutates a caller's data — a frequent, hard-to-trace bug.

## Watch out

- `cp := xs` is NOT a copy of the elements.
- Use `append([]T(nil), xs...)` or `make`+`copy` for independence.
