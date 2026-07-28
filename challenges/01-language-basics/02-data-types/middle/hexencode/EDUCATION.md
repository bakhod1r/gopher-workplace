# Hex encoding

## The idea

A byte is two 4-bit nibbles. Split them and map each through a digit table:

```go
const hexd = "0123456789abcdef"
hi := hexd[b>>4]
lo := hexd[b&0x0f]
```

## Why it matters

Hex encoding is everywhere — hashes, IDs, wire dumps. It exercises shifting,
masking, and table lookup, the same tools as base conversion but fixed at base
16 and fixed width.

## Watch out

- Mask the low nibble with `& 0x0f`; `>>4` already drops the low bits for the
  high one.
- Build with a `[]byte`/`strings.Builder` and convert once — not `+=` per byte.
- Output is fixed two chars per byte, including leading zeros (`0x00` → "00").
