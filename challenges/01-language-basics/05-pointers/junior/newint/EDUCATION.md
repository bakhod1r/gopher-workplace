# Returning pointers to locals

## The idea

Unlike C, Go lets you return `&local`; escape analysis moves it to the heap so it stays valid.

## Why it matters

Constructors routinely return pointers to freshly allocated values.

## Watch out

- Returning `&v` is safe — no dangling pointer.
- `new(int)` also allocates a zeroed int and returns its pointer.
