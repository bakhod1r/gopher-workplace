# RGB Bit Packing

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Packing RGB into `0x00RRGGBB` needs red shifted 16 bits, green 8, blue 0. The
code shifts red by only 8, colliding it with green.

## Task

Fix the single line between the markers in [color.go](color.go) so red occupies
bits 16–23.

## Examples

```go
Pack(0xFF,0,0) // => 0xFF0000
Pack(0,0xFF,0) // => 0x00FF00
Pack(0x12,0x34,0x56) // => 0x123456
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Bit shifting** | Each channel occupies its own 8-bit lane. |
| 2 | **Conversion width** | Convert `uint8` to `uint32` before shifting past 8 bits. |
| 3 | **OR composition** | `|` merges non-overlapping lanes. |

## Hint

Red must shift 16: `uint32(r)<<16 | uint32(g)<<8 | uint32(b)`.

## Validate

```bash
make verify
```
