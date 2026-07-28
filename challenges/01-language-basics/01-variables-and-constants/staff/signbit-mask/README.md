# Sign Bit Mask

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

For an 8-bit two's-complement value the sign is bit 7 — mask `0x80`. The code
masks `0x40` (bit 6), so it misreads `0x40` and `0x80`.

## Task

Fix the single line between the markers in [sign.go](sign.go) so it tests the
correct sign bit.

## Examples

```go
Negative(0x80) // => true  (-128)
Negative(0xFF) // => true  (-1)
Negative(0x40) // => false (64)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Sign bit position** | Bit 7 (`0x80`) for an 8-bit value. |
| 2 | **Bit mask** | `b & 0x80` isolates the top bit. |
| 3 | **Two's complement** | Sign bit set ⇒ negative. |

## Hint

Mask `0x80`.

## Validate

```bash
make verify
```
