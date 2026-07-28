# Pointer values in maps

## The idea

Storing `*int` as map values lets updates flow to the underlying variables, unlike storing ints by value.

## Why it matters

Registries and indexes of mutable objects use pointer values.

## Watch out

- Map iteration order is random but irrelevant here.
- `*p++` changes the referenced variable.
