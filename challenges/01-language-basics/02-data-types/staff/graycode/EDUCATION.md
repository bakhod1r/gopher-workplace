# Reflected Gray code

## Intuition

Binary-to-Gray is a single XOR with the value shifted **right** by one:

```go
x ^ (x >> 1)
```

This guarantees consecutive integers map to codes that differ in exactly one bit.

## Approach

1. Bug: shifts left (x<<1); Gray code uses right shift.
2. Gray = x XOR (x>>1).
3. Fix: return x ^ (x >> 1).

## Solution

```go
func ToGray(x uint32) uint32 {
	return x ^ (x >> 1)
}
```

## Walkthrough

x=4 (100): x>>1=010, 100^010=110=6.

## Pitfalls

- Encoding shifts right; decoding is a different loop (XOR of all higher bits).
- Left vs right shift is the whole bug — direction matters.
- Works per fixed width; the top bit passes through unchanged.
