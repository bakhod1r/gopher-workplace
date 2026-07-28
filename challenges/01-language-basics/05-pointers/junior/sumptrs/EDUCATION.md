# Reading pointer collections safely

## The idea

Dereferencing each non-nil pointer and summing avoids the nil-deref panic on empty slots.

## Why it matters

Aggregating over optional references (sparse arrays) needs nil guards.

## Watch out

- Dereferencing a nil element panics.
- Skip nils, sum the rest.
