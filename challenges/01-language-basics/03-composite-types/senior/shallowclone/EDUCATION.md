# Shallow vs deep copies

## The idea

Copying a struct by value duplicates each field. For a slice field, that copies
the **header** (pointer/len/cap) but not the underlying array — so both structs
share the same elements:

```go
d.Tags = append([]string{}, d.Tags...) // independent copy of the slice
return d
```

## Why it matters

"I made a copy" is only shallow in Go. Reference-typed fields (slices, maps,
pointers) still alias. Mutating the copy then silently corrupts the original — a
frequent, hard-to-trace bug.

## Watch out

- Every slice/map/pointer field needs its own deep copy.
- `slices.Clone`/`maps.Clone` handle one level; nested references need recursion.
- Value types (numbers, strings, arrays) copy fully and are safe.
