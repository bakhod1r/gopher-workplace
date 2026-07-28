# Index and bound must agree

## The idea

`rs := []rune(s)` gives rune-indexed access. The bounds check must use
`len(rs)` (rune count), not `len(s)` (byte count):

```go
if n < 0 || n >= len(rs) { return 0, false }
```

## Why it matters

Mixing units — indexing runes but bounding by bytes — lets invalid indices pass
for ASCII (where byte==rune count) and only fails on non-ASCII text. That is a
classic "works in tests, breaks on real i18n data" bug.

## Watch out

- `len(s)` ≥ `len([]rune(s))`, equal only for pure ASCII.
- The conversion `[]rune(s)` allocates; fine for random access.
- Any time you index one representation, bound with the same one.
