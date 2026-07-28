# Type-asserting a recovered value

## The idea

`recover()` yields `any`; extracting a specific type needs a (comma-ok) assertion — asserting the wrong type silently misses the value.

## Why it matters

Panic values are often errors; asserting the wrong type drops them and reports false success.

## Watch out

- Recovered value is `any`; use `r.(error)` (comma-ok) to get an error.
- A plain assertion of a wrong type would panic — always use comma-ok here.
