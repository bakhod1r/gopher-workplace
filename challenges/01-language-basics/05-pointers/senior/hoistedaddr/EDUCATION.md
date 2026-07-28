# Addresses of reused variables

## The idea

A variable declared outside the loop has one address; taking `&v` each iteration stores the same pointer. A per-iteration declaration gives distinct addresses.

## Why it matters

Even with Go 1.22's per-iteration RANGE vars, a manually hoisted variable reintroduces the shared-address bug.

## Watch out

- `&v` of a hoisted `v` aliases one storage cell.
- Declare the variable inside the loop for distinct addresses.
