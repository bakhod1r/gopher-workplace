# First-class and higher-order functions

## The idea

Functions are ordinary values in Go; a function that takes or returns a function is higher-order, enabling map/filter/reduce.

## Why it matters

It decouples the traversal from the operation, the core of composable data pipelines.

## Watch out

- Return a new slice; don't overwrite the caller's `xs`.
- Preallocating capacity is optional.
