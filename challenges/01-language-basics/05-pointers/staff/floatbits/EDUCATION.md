# Reinterpreting memory with unsafe.Pointer

## Intuition

`*(*T)(unsafe.Pointer(&x))` reads the same bytes as a different type; a value cast instead performs numeric conversion.

## Approach

1. The value is a `float64`; reinterpret its 64 bits directly.
2. The bug narrows to `float32` first, losing precision and reading 32 bits.
3. Cast the address: `*(*uint64)(unsafe.Pointer(&f))`.

## Solution

```go
import "unsafe"

func Bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}
```

## Walkthrough

`float32(f)` rounds and yields a 32-bit pattern, so `Bits(1.5)` comes out wrong. Reading the original `float64` bits gives `0x3ff8000000000000`.

## Pitfalls

- Narrowing to float32 changes the bit pattern and width.
- Reinterpret the float64 directly as uint64 (both 8 bytes).
