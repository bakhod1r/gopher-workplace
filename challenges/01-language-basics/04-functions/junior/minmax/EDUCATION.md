# Multiple return values

## The idea

A function's result list can hold several typed values, returned as a tuple and destructured at the call site.

## Why it matters

Returning related results together (value+error, quotient+remainder, min+max) is idiomatic Go and avoids output parameters.

## Watch out

- Seeding `min` to `0` breaks on all-positive slices.
- Named results are zero-initialised; you still must set them.
