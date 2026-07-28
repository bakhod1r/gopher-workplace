# Reversing Unicode text

## The idea

Reversing the **bytes** of "café" splits the two-byte `é` and produces invalid
UTF-8. Convert to a `[]rune` (one element per character), reverse that, and
convert back:

```go
r := []rune(s)
for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 { r[i], r[j] = r[j], r[i] }
return string(r)
```

## Why it matters

It cements the byte-vs-rune distinction: correct text processing operates on
runes, and `[]rune` gives you random access at the cost of one allocation.

## Watch out

- `[]rune(s)` allocates and decodes; fine here, wasteful in tight loops.
- This reverses code points, not grapheme clusters — combining marks or emoji ZWJ
  sequences can still reorder visibly (a deeper Unicode topic).
- `len([]rune(s))` is the rune count, unlike `len(s)`.
