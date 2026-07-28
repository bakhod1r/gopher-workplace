# Currying with nested closures

## The idea

Each closure captures its argument and returns the next; the innermost sees the whole accumulated environment.

## Why it matters

Currying underlies configurable builders and some functional APIs.

## Watch out

- Every inner closure captures its enclosing arguments by reference.
- The type signature nests exactly as the calls do.
