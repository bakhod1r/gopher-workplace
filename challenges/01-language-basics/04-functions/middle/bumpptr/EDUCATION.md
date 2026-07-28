# Pointers as mutable references

## The idea

Passing `&x` lets a function modify the caller's variable through the pointer; a plain `int` parameter could not.

## Why it matters

In-place mutation APIs (swap, increment, fill) require pointers or slices.

## Watch out

- Forgetting the `*` writes to the pointer variable, not the pointee.
- A nil pointer dereference panics; callers must pass a valid address.
