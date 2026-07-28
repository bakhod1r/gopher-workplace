# Bit rotation

## The idea

Rotating left by `n` moves each bit up `n` places and wraps the top `n` bits to
the bottom. For an 8-bit value:

```go
n &= 7
b<<n | b>>(8-n)
```

The `<<n` moves bits up (the high ones drop off, and `byte` truncates them); the
`>>(8-n)` brings those same high bits back in at the bottom.

## Why it matters

Rotations appear in hashing, cryptography (many ciphers use rotate), and CRC.
Unlike a plain shift, no bits are lost — the pattern is permuted.

## Watch out

- Reduce `n` modulo the width first, or `8-n` goes wrong for `n>=8`.
- With `n==0` after masking, `b>>8` is 0 in Go (shift ≥ width yields 0 for
  unsigned), so the formula still returns `b`.
- Use an unsigned type; rotating signed values sign-extends.
