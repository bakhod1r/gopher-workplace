# Narrowing conversions

## Intuition

Converting to a smaller integer type keeps only the low bits — it **wraps**, it
does not clamp or error:

```go
int8(200)  // -56 (200 mod 256, reinterpreted as signed)
uint8(300) // 44
```

## Approach

1. Apply the explicit conversion int32(n).
2. Go's spec drops the high bits and wraps (two's-complement); no clamping or range check.

## Solution

```go
func ToInt32(n int64) int32 {
	return int32(n)
}
```

## Walkthrough

ToInt32(2147483648): low 32 bits represent int32 min, so int32(2147483648) = -2147483648.

## Pitfalls

- Conversions never panic on overflow; they truncate bits.
- Signedness matters: `int8(200)` reinterprets the high bit as sign.
- Check the range first, or clamp, when the input may exceed the target type.
