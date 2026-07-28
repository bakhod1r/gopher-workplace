# Pointers as references

## The idea

A pointer holds the address of a value; dereferencing (`*p`) reaches the value so a function can mutate the caller's variable.

## Why it matters

In-place mutation (increment, swap, fill) needs a pointer or a slice.

## Watch out

- `p++` moves the pointer variable; `*p++` changes the pointee.
- A nil pointer dereference panics.
