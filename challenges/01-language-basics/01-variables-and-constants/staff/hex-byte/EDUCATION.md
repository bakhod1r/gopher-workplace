# Bit strides when slicing a word

## Intuition

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

## Approach

1. Each byte is 8 bits, so byte `n` is shifted by `8*n`.
2. The bug uses `4*n` (nibble stride).

## Solution

```go
func ByteAt(v uint64, n uint) uint8 {
	return uint8(v >> (8 * n))
}
```

## Walkthrough

`ByteAt(..., 1)` needs `>> 8`; the bug's `>> 4` lands mid-byte. The 8-bit stride selects 0x77.

## Pitfalls

- Stride = field width in bits. Byte → 8, nibble → 4, 16-bit word → 16.
- The truncating conversion keeps only the low bits after the shift; it does not
  save you from a wrong shift amount.
- Shift counts are unsigned and must be `< 64` for a `uint64`.
