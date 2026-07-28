# Mapping a pure transform

## The idea

Producing a new slice from an old one, one transformed element at a time, is the functional map pattern in imperative Go.

## Why it matters

Non-destructive transforms compose cleanly and avoid surprising the caller.

## Watch out

- Preallocating with `make([]int, 0, len(xs))` avoids regrowth but isn't required.
- Clamp each value before appending; don't mutate `xs[i]`.
