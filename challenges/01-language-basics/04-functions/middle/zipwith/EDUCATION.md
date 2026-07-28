# Element-wise combination of two sequences

## The idea

ZipWith iterates two inputs together, bounded by the shorter, applying a two-argument function.

## Why it matters

Vector math, merging parallel columns, and paired transforms use it.

## Watch out

- Stop at the shorter length; indexing past it panics.
- Result length equals the shorter input's length.
