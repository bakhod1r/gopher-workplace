# Pointers to any type

## The idea

A `*bool` mutates a caller's flag just as `*int` mutates an integer.

## Why it matters

Feature flags and dirty bits are toggled through pointers.

## Watch out

- `!*p` reads the current bool then negates.
- Applies to every type, not just numbers.
