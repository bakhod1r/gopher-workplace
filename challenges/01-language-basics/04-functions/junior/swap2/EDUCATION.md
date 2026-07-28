# Returning a tuple

## The idea

A function returning `(int, int)` lets callers do `x, y = Swap(x, y)` — Go evaluates the right side fully before assigning.

## Why it matters

Parallel assignment and multi-return remove the classic temp-variable swap boilerplate.

## Watch out

- The function copies its arguments; it cannot mutate the caller's variables.
- Order in the `return` must match the signature.
