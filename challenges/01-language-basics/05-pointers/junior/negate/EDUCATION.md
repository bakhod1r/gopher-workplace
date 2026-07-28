# Unary mutation through a pointer

## The idea

`*p = -*p` loads, negates, and stores back at the pointee.

## Why it matters

Sign flips and toggles operate in place on shared state.

## Watch out

- Negating twice is a no-op (`-(-x)==x`).
- `-*p` reads first, then assign.
