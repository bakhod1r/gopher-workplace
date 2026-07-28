# Variadic with a validity flag

## The idea

A `...int` plus an `ok bool` cleanly covers the no-argument case without panicking or returning a misleading zero.

## Why it matters

Aggregators that may receive nothing must signal emptiness explicitly.

## Watch out

- Seeding from 0 breaks on all-negative inputs; seed from `nums[0]`.
- Guard before indexing.
