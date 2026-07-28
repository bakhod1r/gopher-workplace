# Pointer arithmetic in bytes

## The idea

`unsafe.Add` works in bytes; indexing into a typed array requires multiplying the index by the element size.

## Why it matters

Manual pointer stepping (interop, custom containers) must account for element stride and alignment.

## Watch out

- `unsafe.Add(base, i)` moves i bytes, not i int32s.
- Multiply by `unsafe.Sizeof(elem)` for element steps.
