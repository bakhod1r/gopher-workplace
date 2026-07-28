# Pointer vs value receivers

## The idea

A value receiver copies the struct; mutations don't reach the caller. Any state-changing method must take a pointer receiver.

## Why it matters

Silent no-op mutation from a value receiver is one of Go's most common bugs.

## Watch out

- `(c Counter)` increments a copy; `(c *Counter)` increments the caller's value.
- Keep receiver types consistent across a type's methods.
