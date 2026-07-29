# Bit packing with shifts

## Intuition

Several small fields can share one integer by giving each its own **bit lane**.
For 8-bit RGB in a `uint32` as `0x00RRGGBB`, red occupies bits 16–23, green 8–15,
blue 0–7:

```go
func Pack(r, g, b uint8) uint32 {
	return uint32(r)<<16 | uint32(g)<<8 | uint32(b)
}
```

Convert each channel to the wider type **before** shifting, then OR the
non-overlapping lanes together. Unpack with a shift-and-mask:

```go
red := uint8(v >> 16)
```

## Approach

1. Red occupies bits 16–23, so shift it by 16, not 8.
2. `r<<16 | g<<8 | b`.

## Solution

```go
func Pack(r, g, b uint8) uint32 {
	return uint32(r)<<16 | uint32(g)<<8 | uint32(b)
}

func Red(v uint32) uint8 { return uint8(v >> 16) }
```

## Walkthrough

The bug shifts red by 8, colliding with green. Shifting by 16 places red in the high byte → 0x123456.

## Pitfalls

- Convert to `uint32` **before** shifting past 8 bits; a `uint8 << 16` is 0.
- Lanes must not overlap — sum the field widths and keep them within the word.
- Mask on unpack (`& 0xFF`) if the field is not at the top.
