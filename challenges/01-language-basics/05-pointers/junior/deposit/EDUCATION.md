# Pointer receivers by hand

## The idea

A function taking `*Account` and updating a field is the mechanism behind pointer-receiver methods.

## Why it matters

Stateful domain objects (accounts, carts) mutate through pointers.

## Watch out

- A value parameter would update a copy, losing the change.
- `a.Balance += amount` auto-dereferences the pointer.
