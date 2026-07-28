# Read-modify-write through a pointer

## The idea

`*p += d` loads, adds, and stores at the pointee in one expression.

## Why it matters

Shared counters and accumulators updated by helpers use this pattern.

## Watch out

- `*p += delta` mutates the caller's variable.
- Works for negative deltas too.
