# Bytes vs runes

## The idea

A Go string is a read-only sequence of **bytes** holding UTF-8. `s[i]` indexes a
byte. A **rune** is a Unicode code point (an `int32`), and one rune may occupy
1–4 bytes. Ranging over a string decodes UTF-8 and yields runes:

```go
for _, r := range s { return r } // first rune, whole character
```

## Why it matters

`s[0]` for `"étage"` is the first *byte* of `é` (0xC3), not the character. Any
text handling that assumes one byte = one character breaks on non-ASCII. The
`for range` form is the correct way to walk characters.

## Watch out

- The range index is a **byte offset**, so it jumps by rune width (0, then the
  byte length of each rune).
- `len(s)` is the byte count, not the rune count; use `utf8.RuneCountInString`.
- Converting `[]rune(s)` gives random access to runes at the cost of a copy.

## Try it yourself

```go
for i, r := range "é!" { fmt.Println(i, r) } // 0 233 ; 2 33
len("é")                                     // 2 (bytes)
[]rune("é")[0]                               // 233
```
