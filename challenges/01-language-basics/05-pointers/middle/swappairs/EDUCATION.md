# Pointer surgery on adjacent nodes

## The idea

Swapping nodes rewires their Next links and promotes the second node to head — order the assignments so no link is lost.

## Why it matters

Pairwise swaps generalise to the full 'swap nodes in pairs' problem.

## Watch out

- Save `second.Next` before overwriting.
- Return the new head (the old second node).
