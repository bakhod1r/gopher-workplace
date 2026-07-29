# Width matching in reinterprets

## Intuition

A reinterpret cast must use a target type the same size as the source; a narrower type reads only part of the memory.

## Approach

1. The value is 64-bit; read 64 bits.
2. The bug reads a `uint32`, dropping the high half.
3. `*(*uint64)(unsafe.Pointer(&x))`.

## Solution

```go
import "unsafe"

func AsU64(x int64) uint64 {
	return *(*uint64)(unsafe.Pointer(&x))
}
```

## Walkthrough

Reading only 32 bits of `-1` yields `0xffffffff`. The full 64-bit reinterpret gives `0xffffffffffffffff`.

## Pitfalls

- int64 is 8 bytes; `*uint32` reads only the low 4.
- Match widths: `*uint64` for an int64.
