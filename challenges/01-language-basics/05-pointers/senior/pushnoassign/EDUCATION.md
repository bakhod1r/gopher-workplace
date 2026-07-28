# Reassigning appended slice fields

## The idea

A struct's slice field must be updated with the append result; ignoring it keeps the old, shorter header.

## Why it matters

Dropping the append result is a classic silent no-op growth bug.

## Watch out

- `_ = append(s.data, v)` builds nothing observable.
- Write `s.data = append(s.data, v)`.
