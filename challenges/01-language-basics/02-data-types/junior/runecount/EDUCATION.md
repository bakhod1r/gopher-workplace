# Counting runes, not bytes

## The idea

`len(s)` is the **byte** count. The number of characters is the rune count:

```go
utf8.RuneCountInString(s) // characters
len(s)                    // bytes (>= rune count)
```

Ranging over a string also yields runes, so counting iterations works too.

## Why it matters

Length limits, column widths, and "how many characters" all mean runes for
human-facing text. Using `len(s)` overcounts any non-ASCII string.

## Watch out

- `len(s)` equals the rune count only for pure ASCII.
- `for range s` decodes UTF-8; the index is a byte offset that jumps by rune
  width.
- A rune is a code point, not necessarily a full grapheme (combining marks).
