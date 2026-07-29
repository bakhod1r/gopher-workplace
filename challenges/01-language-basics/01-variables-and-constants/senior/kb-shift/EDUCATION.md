# Shift stride in iota units

## Intuition

Binary size units scale by 1024 per step, which is a **10-bit** shift, not one
bit. The `iota` slot picks the exponent, so the multiplier is `10 * iota`:

```go
const (
	_  ByteSize = iota          // consumes iota==0
	KB ByteSize = 1 << (10 * iota) // iota==1 -> 1<<10 == 1024
	MB                              // iota==2 -> 1<<20
	GB                              // iota==3 -> 1<<30
)
```

`1 << iota` (stride 1) would give 2, 4, 8 — powers of two, not powers of 1024.

## Approach

1. Each unit is `1 << (10 * iota)`, a 1024× step.
2. The bug uses `1 << iota` (2× step); restore the `10 *`.

## Solution

```go
type ByteSize uint64

const (
	_ ByteSize = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
)
```

## Walkthrough

The bug makes KB `1 << 1 = 2`; `1 << (10 * iota)` gives `1 << 10 = 1024`.

## Pitfalls

- The blank `_ = iota` line matters: it puts KB on `iota==1`, so `10*iota == 10`.
- Use `uint64`; `1 << 40` (TB) overflows 32-bit ints.
- Verify by pinning KB to exactly 1024 in a test.
