# Pointer identity vs pointed-to equality

## The idea

`a == b` tests whether two pointers share an address; `*a == *b` tests whether their values match — different questions.

## Why it matters

Confusing the two breaks caches, sets, and cycle checks that rely on identity.

## Watch out

- `*a == *b` is true for two different vars holding 5.
- Use `a == b` for identity.
