# Reversing byte order

## Intuition

Each byte moves to the mirrored lane, and the mask must match the **destination**
lane:

```go
x<<24                 // byte0 -> byte3 (no mask; shifted out bits vanish)
(x<<8)  & 0xFF0000    // byte1 -> byte2
(x>>8)  & 0xFF00      // byte2 -> byte1
x>>24                 // byte3 -> byte0
```

## Approach

1. Bug: second lane mask is 0xFF00 instead of 0xFF0000, corrupting byte 2.
2. Reversing 4 bytes needs lanes masked at 0xFF000000/0xFF0000/0xFF00/0xFF.
3. Fix: (x<<8)&0xFF0000.

## Solution

```go
func Swap32(x uint32) uint32 {
	return x<<24 | (x<<8)&0xFF0000 | (x>>8)&0xFF00 | x>>24
}
```

## Walkthrough

0x11223344: <<24=0x44..., (x<<8)&0xFF0000=0x330000, (x>>8)&0xFF00=0x2200, >>24=0x11 -> 0x44332211.

## Pitfalls

- The mask names where a byte **lands**, after the shift.
- Outer lanes (`<<24`, `>>24`) need no mask; the shift discards the rest.
- `math/bits.ReverseBytes32` does this in one instruction.
