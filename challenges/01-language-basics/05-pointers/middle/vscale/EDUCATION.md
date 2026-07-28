# Value receivers and immutability

## The idea

A value-receiver method works on a copy; returning a new value keeps the original immutable — natural for small value types.

## Why it matters

Immutable value semantics (points, complex numbers, money) use value receivers that return new values.

## Watch out

- Mutating `p` inside a value receiver changes only the copy.
- Return a new struct to express a transform.
