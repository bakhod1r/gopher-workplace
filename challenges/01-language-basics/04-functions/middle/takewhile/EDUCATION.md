# Prefix-based higher-order operations

## The idea

TakeWhile/DropWhile depend on order and stop at the boundary, unlike Filter which scans everything.

## Why it matters

Parsing and streaming logic often consumes matching prefixes.

## Watch out

- `break` on the first failure — do not continue scanning like Filter.
- Empty result when the first element already fails.
