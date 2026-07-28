# Bit strides when slicing a word

## The idea

To pull out field `n` you shift by `n * width`. A byte is 8 bits wide, so byte
`n` sits at bit offset `8 * n`:

```go
func ByteAt(v uint64, n uint) uint8 {
	return uint8(v >> (8 * n))
}
```

A stride of 4 (a nibble) reads a value straddling two bytes; the final
`uint8(...)` truncation then hides the error for `n == 0`, where both strides
agree.

## Why it matters

Serialization, hashing, colour, and network code slice words into fixed fields.
The wrong stride corrupts every field except the zeroth, and the zeroth still
looks correct — a debugging trap.

## Watch out

- Stride = field width in bits. Byte → 8, nibble → 4, 16-bit word → 16.
- The truncating conversion keeps only the low bits after the shift; it does not
  save you from a wrong shift amount.
- Shift counts are unsigned and must be `< 64` for a `uint64`.

## Try it yourself

```go
v := uint64(0xAABBCCDD)
uint8(v >> (8 * 0)) // 0xDD
uint8(v >> (8 * 1)) // 0xCC
uint8(v >> (8 * 3)) // 0xAA
```
