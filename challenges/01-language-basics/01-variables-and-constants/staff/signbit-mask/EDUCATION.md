# The sign bit

## Intuition

In an N-bit two's-complement value the **top** bit (bit N-1) is the sign: set
means negative. For 8 bits that is bit 7, mask `0x80`:

```go
func Negative(b uint8) bool { return b&0x80 != 0 }
```

Masking `0x40` (bit 6) tests the wrong bit, so it misreads exactly the values
where bit 7 and bit 6 differ — including `0x80` (-128) and `0x40` (64).

## Approach

1. The sign bit of a byte is bit 7 → mask `0x80`.
2. The bug masks `0x40` (bit 6).

## Solution

```go
func Negative(b uint8) bool {
	return b&0x80 != 0
}
```

## Walkthrough

`Negative(0x80)` needs bit 7; `& 0x40` misses it. `& 0x80` correctly reports the sign bit.

## Pitfalls

- Sign bit position is width-dependent: `0x80` for 8-bit, `0x8000` for 16-bit,
  `1<<63` for 64-bit.
- `b & mask != 0` tests presence; do not compare `== mask` unless only that bit
  may be set.
- Converting `uint8` to `int8` also reveals the sign, if you prefer that route.
