# Output parameters via pointers

## The idea

Writing results through caller-supplied pointers is an alternative to multiple return values, common when interfacing with fixed signatures.

## Why it matters

Fill-style APIs (Scan, Decode) populate caller variables through pointers.

## Watch out

- Dereference to write the results: `*min = v`.
- Go usually prefers multiple returns, but pointer outputs appear in stdlib.
