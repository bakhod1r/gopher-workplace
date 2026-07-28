# Optional values as pointers

## The idea

A `*int` doubles as an optional int; the nil case maps to the zero value.

## Why it matters

Config and JSON decoding use pointers to distinguish 'unset' from zero.

## Watch out

- Reading `*p` before the nil check panics.
- Returning 0 collapses nil and a real 0 — fine when that's intended.
