# Iterative list reversal

## The idea

Walking the list while re-pointing each node's Next to the previous node reverses it in O(n) time and O(1) space.

## Why it matters

List reversal is a canonical pointer-manipulation exercise.

## Watch out

- Save `cur.Next` before overwriting it, or you lose the rest.
- The final `prev` is the new head.
