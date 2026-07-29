# Counting set bits

## Intuition

`x & (x-1)` clears the **lowest** set bit of `x`. Repeat until zero and you have
counted exactly the set bits:

```go
for x != 0 { x &= x - 1; n++ }
```

Subtracting 1 flips the lowest 1 to 0 and all bits below it to 1; ANDing with
the original keeps only the higher bits.

## Approach

1. Loop while x!=0. 2. Clear the lowest set bit with x &= x-1. 3. Count each clear. The loop runs once per set bit.

## Solution

```go
func Count(x uint64) int {
	n := 0
	for x != 0 {
		x &= x - 1
		n++
	}
	return n
}
```

## Walkthrough

Count(0b1011): 1011->1010->1000->0000, 3 iterations -> 3.

## Pitfalls

- Use an unsigned type so the right shift alternative is logical, not
  arithmetic.
- The stdlib `math/bits.OnesCount64` does this in one instruction — hand-rolling
  is for understanding.
- `x-1` on `x==0` underflows to all-ones, but the loop guard `x != 0` prevents
  entering with 0.
