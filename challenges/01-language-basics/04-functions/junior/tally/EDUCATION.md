# Computing multiple aggregates in one pass

## The idea

Accumulating several results in a single loop avoids re-scanning and keeps related outputs in sync.

## Why it matters

Stats functions (count, sum, min, max) commonly return a bundle from one traversal.

## Watch out

- `n%2 != 0` for negative odds too (`-3%2 == -1`); evenness test still holds.
- Both accumulators start at 0.
