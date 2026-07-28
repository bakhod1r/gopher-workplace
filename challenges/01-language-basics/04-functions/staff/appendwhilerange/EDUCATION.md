# Bounding loops that grow their own target

## The idea

A loop appending to `out` must not use `len(out)` as its bound; iterate the fixed source length instead to avoid skips or runaway growth.

## Why it matters

Self-referential loop bounds over a mutating slice cause skipped work or infinite growth.

## Watch out

- Bound on the stable input (`len(xs)`), never the growing output.
- `for i := range xs` snapshots the count safely.
