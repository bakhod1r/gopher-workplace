# Pointers to arrays

## The idea

Arrays are values; a `*[N]T` lets a function mutate the caller's array, and indexing auto-dereferences the pointer.

## Why it matters

Fixed-size buffers passed by pointer avoid copying and allow in-place edits.

## Watch out

- Passing `[4]int` by value copies; the caller wouldn't see changes.
- `arr[i]` on a pointer is `(*arr)[i]`.
