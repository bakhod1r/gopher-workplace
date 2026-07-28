# Reflected Gray code

## The idea

Binary-to-Gray is a single XOR with the value shifted **right** by one:

```go
x ^ (x >> 1)
```

This guarantees consecutive integers map to codes that differ in exactly one bit.

## Why it matters

Rotary/position encoders and some error-tolerant schemes rely on the one-bit
property to avoid transient glitches when multiple bits would flip at once. A
left shift produces a different, non-Gray sequence that loses the property.

## Watch out

- Encoding shifts right; decoding is a different loop (XOR of all higher bits).
- Left vs right shift is the whole bug — direction matters.
- Works per fixed width; the top bit passes through unchanged.
