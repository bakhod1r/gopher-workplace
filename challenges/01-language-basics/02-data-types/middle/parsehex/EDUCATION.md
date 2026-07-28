# Parsing hex

## The idea

Each hex digit contributes 4 bits: `n = n*16 + d`. The digit value `d` depends on
which of three contiguous ranges the character falls in:

```go
switch {
case c >= '0' && c <= '9': d = int(c - '0')
case c >= 'a' && c <= 'f': d = int(c-'a') + 10
case c >= 'A' && c <= 'F': d = int(c-'A') + 10
default: return 0, false
}
```

## Why it matters

It is the inverse of hex encoding and generalizes `atoi` to base 16 with
case-folding — a compact exercise in character-range classification.

## Watch out

- Handle both letter cases.
- `n*16` can overflow for long inputs; this simplified version ignores that.
- Reject empty input explicitly, or it returns `(0, true)`.
