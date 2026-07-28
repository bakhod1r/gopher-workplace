# Base cases in recursion

## The idea

A recursive definition needs a base that both terminates and returns the identity for the combining operation — 1 for products, 0 for sums.

## Why it matters

An off base value doesn't crash; it silently returns wrong results.

## Watch out

- `0! == 1` by definition; the multiplicative identity is 1.
- Wrong base values are hard to spot because the recursion still terminates.
