# unsafe.Slice length semantics

## The idea

`unsafe.Slice(ptr, n)` builds a slice of n ELEMENTS; passing a byte length creates an out-of-bounds view.

## Why it matters

Constructing slices over C or mmap'd memory must use element counts, not byte sizes.

## Watch out

- `unsafe.Slice`'s length is in elements, not bytes.
- An over-long length yields a slice that reads past the array.
