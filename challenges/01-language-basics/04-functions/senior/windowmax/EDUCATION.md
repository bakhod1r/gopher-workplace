# Sliding-window index arithmetic

## The idea

When the window covers `[i-k+1 .. i]` after adding `xs[i]`, the element that just left is `xs[i-k]`; off-by-one here is a classic bug.

## Why it matters

Sliding windows power streaming stats, rate limiting, and substring problems; the boundary index must be exact.

## Watch out

- Adding `xs[i]` means removing `xs[i-k]` to keep width k.
- Draw the indices for one step to check the boundary.
