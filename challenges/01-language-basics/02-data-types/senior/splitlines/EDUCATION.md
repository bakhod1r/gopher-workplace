# Normalizing line endings

## The idea

Split on `\n`, then strip a trailing `\r` from each piece so CRLF and LF inputs
produce identical lines:

```go
parts[i] = strings.TrimSuffix(p, "\r")
```

## Why it matters

Text arrives from many platforms. A stray `\r` is invisible in output but breaks
equality checks, map keys, and numeric parses. Normalizing at the split is the
robust fix.

## Watch out

- `TrimSuffix` removes `\r` only if present — safe for LF-only input.
- Don't `TrimSpace`; that would also drop meaningful leading/trailing spaces.
- Some files end with a trailing newline, yielding a final empty line — decide
  whether to keep it.
