# Storing computed results through a pointer

## The idea

Calling a pure function without assigning its result is a no-op; the transformed value must be written back through the pointer.

## Why it matters

Discarding a computed result is a classic silent no-op bug.

## Watch out

- `f(*p)` returns a value nobody keeps.
- `*p = f(*p)` updates the caller's variable.
