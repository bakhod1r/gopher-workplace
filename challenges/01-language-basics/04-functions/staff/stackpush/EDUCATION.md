# Recording per-iteration data in defers

## The idea

A deferred closure that reads a fixed expression captures none of the loop's progression; snapshot `xs[i]` per iteration to record distinct values.

## Why it matters

Deferred accumulation that reads a constant index is a real logic bug masked by LIFO ordering.

## Watch out

- `xs[len(xs)-1]` is the same element every iteration.
- Snapshot `xs[i]` as an argument for a version-proof reverse.
