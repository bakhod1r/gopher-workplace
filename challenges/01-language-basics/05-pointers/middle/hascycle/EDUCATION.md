# Floyd's cycle detection

## The idea

Two pointers at speeds 1 and 2 must meet inside a cycle and never meet in an acyclic list — O(1) space.

## Why it matters

Cycle detection guards against infinite loops in linked structures.

## Watch out

- Guard `fast` and `fast.Next` before the double step.
- Reaching nil proves no cycle.
