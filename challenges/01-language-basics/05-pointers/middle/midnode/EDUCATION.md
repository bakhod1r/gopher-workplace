# The tortoise-and-hare technique

## The idea

Two pointers at different speeds locate the middle (or detect cycles) in a single pass with O(1) extra space.

## Why it matters

Slow/fast pointers power cycle detection, midpoint finding, and Nth-from-end.

## Watch out

- Guard both `fast != nil` and `fast.Next != nil` before advancing.
- The even-length midpoint convention depends on the loop condition.
