# Prefix-sum range queries

## The idea

With `pre[i]` the sum of the first i elements, any range `[l:r)` is `pre[r] - pre[l]`; dropping the `pre[l]` term over-counts.

## Why it matters

Prefix sums turn O(n) range queries into O(1); the subtraction is the whole point.

## Watch out

- `pre[r]` alone is the sum from 0, not from l.
- Match the half-open convention: `pre[r] - pre[l]` for `xs[l:r]`.
