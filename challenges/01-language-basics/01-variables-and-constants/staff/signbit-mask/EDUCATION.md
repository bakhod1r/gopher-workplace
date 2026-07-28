# The sign bit

## The idea

In an N-bit two's-complement value the **top** bit (bit N-1) is the sign: set
means negative. For 8 bits that is bit 7, mask `0x80`:

```go
func Negative(b uint8) bool { return b&0x80 != 0 }
```

Masking `0x40` (bit 6) tests the wrong bit, so it misreads exactly the values
where bit 7 and bit 6 differ — including `0x80` (-128) and `0x40` (64).

## Why it matters

Sign detection on raw bytes shows up in decoders, VM opcodes, and compression.
An off-by-one bit mask is easy to write and passes any test that happens not to
exercise the boundary between the two bits.

## Watch out

- Sign bit position is width-dependent: `0x80` for 8-bit, `0x8000` for 16-bit,
  `1<<63` for 64-bit.
- `b & mask != 0` tests presence; do not compare `== mask` unless only that bit
  may be set.
- Converting `uint8` to `int8` also reveals the sign, if you prefer that route.

## Try it yourself

```go
uint8(0x80) & 0x80 // 0x80 -> negative
uint8(0x7F) & 0x80 // 0    -> non-negative
int8(0x80)         // -128
```
