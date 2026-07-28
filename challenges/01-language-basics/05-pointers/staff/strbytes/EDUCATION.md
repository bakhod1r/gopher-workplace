# String header fields

## The idea

A string is a two-word header: a data pointer and an int length. Reading the wrong word returns an address instead of the length.

## Why it matters

Manual header access must target the correct field; mixing them yields garbage.

## Watch out

- The first word is the data pointer; the second is the length.
- `len(s)` is the safe, idiomatic way; this exercises the header layout.
