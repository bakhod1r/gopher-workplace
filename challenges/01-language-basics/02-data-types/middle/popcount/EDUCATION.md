# Counting set bits

## The idea

`x & (x-1)` clears the **lowest** set bit of `x`. Repeat until zero and you have
counted exactly the set bits:

```go
for x != 0 { x &= x - 1; n++ }
```

Subtracting 1 flips the lowest 1 to 0 and all bits below it to 1; ANDing with
the original keeps only the higher bits.

## Why it matters

Popcount powers Hamming distance, sparse-set sizes, and bitboard tricks. The
`x&(x-1)` idiom runs in as many steps as there are 1-bits, faster than scanning
all 64.

## Watch out

- Use an unsigned type so the right shift alternative is logical, not
  arithmetic.
- The stdlib `math/bits.OnesCount64` does this in one instruction — hand-rolling
  is for understanding.
- `x-1` on `x==0` underflows to all-ones, but the loop guard `x != 0` prevents
  entering with 0.
