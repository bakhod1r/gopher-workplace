# Constructing struct pointers

## The idea

`&T{...}` allocates and initialises in one expression; it's the idiomatic constructor, equivalent to `new(T)` followed by field sets.

## Why it matters

Factory functions return `*T` built with a composite literal.

## Watch out

- `&User{...}` is clearer than `new(User)` + assignments.
- Each call allocates a distinct instance.
