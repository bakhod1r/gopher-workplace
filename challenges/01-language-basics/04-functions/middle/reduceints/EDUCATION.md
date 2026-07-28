# Folding with an accumulator

## The idea

Reduce expresses sum, product, max, and concatenation as one shape parameterised by the combining function and seed.

## Why it matters

It is the general form behind most aggregate computations.

## Watch out

- With an empty slice, the result is `init`.
- Argument order in `f(acc, x)` must match the signature.
