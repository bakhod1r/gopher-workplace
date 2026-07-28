# Bounded character ranges

## The idea

ASCII uppercase letters occupy 65..90. Lowercasing adds 32 — but only for that
exact range. A one-sided check (`c >= 'A'`) also matches `[`, `\`, `]`, `^`, `_`,
backtick, and every lowercase letter, corrupting them:

```go
if c >= 'A' && c <= 'Z' { b[i] = c + 32 }
```

## Why it matters

Case-folding headers, tokens, and identifiers must touch letters only. A missing
upper bound is invisible on pure-letter tests and silently mangles punctuation
and symbols in real input.

## Watch out

- Always bound both ends of a character range.
- `unicode.ToLower` handles non-ASCII correctly; this ASCII-only version is for
  bytes.
- The `+32` offset is specific to ASCII letters.
