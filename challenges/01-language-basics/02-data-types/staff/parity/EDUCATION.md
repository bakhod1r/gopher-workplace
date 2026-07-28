# Parity via XOR

## The idea

The parity bit is 1 exactly when an odd number of bits are set. XOR is the parity
of two bits, so folding every bit with XOR yields the parity of the whole word:

```go
p ^= int(x & 1)
```

## Why it matters

Parity is the simplest error-detection code (serial links, RAID, ECC building
block). Summing set bits returns a **population count**, not a single parity bit —
a category error that fails the moment more than one bit is set.

## Watch out

- `^=` folds to 0/1; `+=` accumulates the count.
- `math/bits.OnesCount32(x) & 1` is the same parity in one step.
- Parity detects an odd number of bit errors, not their location.
