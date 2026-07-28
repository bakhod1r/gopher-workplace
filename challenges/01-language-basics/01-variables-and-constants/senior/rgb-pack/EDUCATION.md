# Bit packing with shifts

## The idea

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

## Why it matters

Packing is everywhere: colours, flags, network headers, compact keys. If two
fields share a shift amount (e.g. red also shifted 8), their bits collide and the
value is corrupt in a way that is invisible until you unpack.

## Watch out

- Convert to `uint32` **before** shifting past 8 bits; a `uint8 << 16` is 0.
- Lanes must not overlap — sum the field widths and keep them within the word.
- Mask on unpack (`& 0xFF`) if the field is not at the top.

## Try it yourself

```go
pack := func(hi, lo uint8) uint16 { return uint16(hi)<<8 | uint16(lo) }
pack(0x12, 0x34) // 0x1234
uint8(0x1234 >> 8) // 0x12
```
