# The delete idiom

## The idea

Deleting index `i` keeps everything before it and everything **after** it:

```go
append(xs[:i], xs[i+1:]...)
```

Using `xs[i:]` for the tail re-includes the element you meant to drop.

## Why it matters

Off-by-one in the tail bound is a common slice-delete bug; the length is right but
the target survives (and the last element is duplicated/lost). `slices.Delete`
encapsulates the correct bounds.

## Watch out

- Tail starts at `i+1`, not `i`.
- This overwrites the backing array — copy first if the caller's slice matters.
- Guard `i` in range before slicing.
