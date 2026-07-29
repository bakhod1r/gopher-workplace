# Parity via XOR

## Intuition

The parity bit is 1 exactly when an odd number of bits are set. XOR is the parity
of two bits, so folding every bit with XOR yields the parity of the whole word:

```go
p ^= int(x & 1)
```

## Approach

1. Bug: `p += bit` accumulates the popcount, not parity.
2. Parity is the XOR of all bits, giving 0 or 1.
3. Fix: p ^= int(x & 1).

## Solution

```go
func Parity(x uint32) int {
	p := 0
	for x != 0 {
		p ^= int(x & 1)
		x >>= 1
	}
	return p
}
```

## Walkthrough

x=3: bits 1,1 -> 0^1^1 = 0.

## Pitfalls

- `^=` folds to 0/1; `+=` accumulates the count.
- `math/bits.OnesCount32(x) & 1` is the same parity in one step.
- Parity detects an odd number of bit errors, not their location.
