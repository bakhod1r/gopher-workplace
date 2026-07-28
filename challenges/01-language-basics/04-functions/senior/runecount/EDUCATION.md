# String length in bytes vs runes

## The idea

A Go string is UTF-8 bytes; `len` returns bytes while ranging or `utf8.RuneCountInString` returns code points.

## Why it matters

Byte/rune confusion corrupts indexing, truncation, and length limits for non-ASCII text.

## Watch out

- `len("é")` is 2, not 1.
- Range indices over a string are BYTE offsets, but the value is a rune.
