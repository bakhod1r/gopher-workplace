# LEB128 varints

## The idea

A varint stores 7 value bits per byte, low group first; the 8th bit signals
whether more bytes follow. Each successive group shifts up by **7**, not 8:

```go
v |= uint64(c&0x7F) << shift
if c&0x80 == 0 { return v, i + 1 }
shift += 7
```

## Why it matters

Varints are the backbone of protobuf, DWARF, and many binary formats. Advancing
by 8 leaves a zero gap every byte, so any multi-byte value decodes wrong — a
subtle, format-breaking bug.

## Watch out

- Mask the payload with `0x7F`; the top bit is the continuation flag, not data.
- Groups are little-endian: the first byte is the least significant 7 bits.
- Guard against a slice that ends with the continuation bit still set.
