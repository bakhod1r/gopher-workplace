# In-place filtering aliases the input

## The idea

`xs[:0]` is a length-0 slice over `xs`'s **same** backing array. Appending to it
writes over `xs[0], xs[1], ...`. That's the deliberate "filter in place" trick —
but only safe when you *want* to overwrite the input:

```go
out := []int{} // independent result; input untouched
```

## Why it matters

The `xs[:0]` idiom is powerful but sharp: used accidentally, it corrupts the
caller's data as it filters. Knowing when reuse is intended vs a bug is the
lesson.

## Watch out

- `xs[:0]` shares memory; a fresh slice does not.
- The in-place version (used by `slices.DeleteFunc`) mutates and returns the
  input.
- Reading `xs` after an in-place filter over it is undefined-ish — the front is
  overwritten.
