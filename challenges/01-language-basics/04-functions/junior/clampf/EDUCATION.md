# Bounding a value

## The idea

Clamping composes a lower and upper limit into one operation used across UI, graphics, and numeric guards.

## Why it matters

Prevents out-of-range indices, colors, or volumes from propagating downstream.

## Watch out

- With `lo > hi` the range is empty; the task assumes `lo <= hi`.
- Use inclusive comparisons so the endpoints are reachable.
