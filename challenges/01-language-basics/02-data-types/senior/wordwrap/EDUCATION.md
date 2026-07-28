# Inclusive width boundaries

## The idea

A word fits on the current line when the resulting length — existing text, one
space, and the word — is **at most** the width. The boundary is inclusive:

```go
if len(line)+1+len(w) <= width { line += " " + w }
```

Using `<` rejects a line that exactly fills the width, wrapping one character too
soon.

## Why it matters

Layout, pagination, and column fitting hinge on whether the boundary is inclusive.
Off-by-one here produces ragged output and wasted space in every UI that wraps
text.

## Watch out

- Include the separator (`+1`) in the length math.
- `<=` vs `<` is the whole bug — inclusive width.
- This simplified version counts bytes; real wrapping counts runes/display width.
