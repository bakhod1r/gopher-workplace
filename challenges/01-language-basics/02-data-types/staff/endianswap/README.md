# Byte-Order Swap

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A network layer converts a 32-bit field between host and network byte order. One
lane uses the mask `0xFF00` where it should be `0xFF0000`, so byte 1 lands in the
wrong place.

## Task

Fix the masks between the markers in [endianswap.go](endianswap.go).

## Examples

```go
Swap32(0x11223344) // => 0x44332211
Swap32(0xDEADBEEF) // => 0xEFBEADDE
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Byte lanes** | Bytes occupy bits 0-7, 8-15, 16-23, 24-31. |
| 2 | **Shift + mask** | Move each byte, mask its destination lane. |
| 3 | **Symmetry** | Outer bytes need no mask; inner two do. |

## Hint

The `(x<<8)` term must mask `0xFF0000`: `x<<24 | (x<<8)&0xFF0000 | (x>>8)&0xFF00 | x>>24`.

## Validate

```bash
make verify
```
