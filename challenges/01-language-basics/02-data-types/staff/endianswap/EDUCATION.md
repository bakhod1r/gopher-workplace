# Reversing byte order

## The idea

Each byte moves to the mirrored lane, and the mask must match the **destination**
lane:

```go
x<<24                 // byte0 -> byte3 (no mask; shifted out bits vanish)
(x<<8)  & 0xFF0000    // byte1 -> byte2
(x>>8)  & 0xFF00      // byte2 -> byte1
x>>24                 // byte3 -> byte0
```

## Why it matters

Endian conversion is mandatory at every network/file boundary. A wrong mask
misplaces one byte, and the corruption is invisible for symmetric values (like
`0x11223344` reversed is still plausible) — it bites on real data.

## Watch out

- The mask names where a byte **lands**, after the shift.
- Outer lanes (`<<24`, `>>24`) need no mask; the shift discards the rest.
- `math/bits.ReverseBytes32` does this in one instruction.
