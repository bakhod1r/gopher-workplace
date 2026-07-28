# Shared vs per-closure captured cells

## The idea

Closures capture variables by identity; a variable declared outside the loop is one cell shared by all, while one declared inside gives each closure its own.

## Why it matters

Sharing a single accumulator across supposedly-independent closures is a classic state-leak bug.

## Watch out

- Hoisting `c` out of the loop makes every closure share it.
- Declare per-iteration state inside the loop body.
