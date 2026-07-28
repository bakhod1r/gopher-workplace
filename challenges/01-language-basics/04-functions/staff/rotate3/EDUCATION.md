# Multiple assignment evaluation order

## The idea

Go evaluates all right-hand-side operands before assigning any left-hand target, enabling rotations and swaps without temporaries.

## Why it matters

Emulating parallel assignment with sequential statements is a classic clobbering bug.

## Watch out

- `a, b, c = b, c, a` moves all three at once.
- Sequential `a = b; ...; c = a` uses the already-changed `a`.
