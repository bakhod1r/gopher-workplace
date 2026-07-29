# LEB128 varints

## Intuition

A varint stores 7 value bits per byte, low group first; the 8th bit signals
whether more bytes follow. Each successive group shifts up by **7**, not 8:

```go
v |= uint64(c&0x7F) << shift
if c&0x80 == 0 { return v, i + 1 }
shift += 7
```

## Approach

1. Bug: shift += 8; LEB128 packs 7 payload bits per byte.
2. Each byte contributes 7 bits, so the shift must grow by 7.
3. Fix: shift += 7.

## Solution

```go
func Decode(b []byte) (uint64, int) {
	var v uint64
	var shift uint
	for i := 0; i < len(b); i++ {
		c := b[i]
		v |= uint64(c&0x7F) << shift
		if c&0x80 == 0 {
			return v, i + 1
		}
		shift += 7
	}
	return 0, 0
}
```

## Walkthrough

[0xAC,0x02]: 0x2C=44 at shift0, 0x02=2 at shift7 -> 44 | 256 = 300.

## Pitfalls

- Mask the payload with `0x7F`; the top bit is the continuation flag, not data.
- Groups are little-endian: the first byte is the least significant 7 bits.
- Guard against a slice that ends with the continuation bit still set.
