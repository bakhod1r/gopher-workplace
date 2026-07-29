# Bit rotation

## Intuition

Rotating left by `n` moves each bit up `n` places and wraps the top `n` bits to
the bottom. For an 8-bit value:

```go
n &= 7
b<<n | b>>(8-n)
```

The `<<n` moves bits up (the high ones drop off, and `byte` truncates them); the
`>>(8-n)` brings those same high bits back in at the bottom.

## Approach

1. Mask n to 0..7 with n&=7. 2. If n==0 return b unchanged (avoids shifting by 8). 3. Otherwise b<<n | b>>(8-n) recombines the wrapped bits.

## Solution

```go
func Left(b byte, n int) byte {
	n &= 7
	if n == 0 {
		return b
	}
	return b<<n | b>>(8-n)
}
```

## Walkthrough

Left(0b1000_0000,1): n=1. b<<1=0 (top bit off), b>>7=1. 0|1 = 0b0000_0001.

## Pitfalls

- Reduce `n` modulo the width first, or `8-n` goes wrong for `n>=8`.
- With `n==0` after masking, `b>>8` is 0 in Go (shift ≥ width yields 0 for
  unsigned), so the formula still returns `b`.
- Use an unsigned type; rotating signed values sign-extends.
