# Collections of pointers

## The idea

A `[]*int` holds addresses; iterating and dereferencing mutates each target, with nil entries needing a guard.

## Why it matters

Pointer slices model editable collections and observer lists.

## Watch out

- Dereferencing a nil element panics — skip it.
- Mutating `*p` changes the original variable.
