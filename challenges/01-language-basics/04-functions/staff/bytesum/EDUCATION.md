# Choosing an accumulator type

## The idea

Summation can exceed the element type's range; the accumulator must be sized for the aggregate, or unsigned/overflow wraps the result silently.

## Why it matters

Narrow accumulators over wide data are a real, silent numeric-overflow bug.

## Watch out

- `uint8 += ...` wraps at 256; the sum of small bytes can still overflow.
- Accumulate in `int` (or `uint64`) and convert the element, not the total.
