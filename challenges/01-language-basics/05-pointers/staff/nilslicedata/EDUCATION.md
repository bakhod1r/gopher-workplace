# Data pointers of empty slices

## The idea

`unsafe.SliceData` of an empty slice can be nil; you must check the length before dereferencing it, just like any pointer.

## Why it matters

Skipping the length guard on the data pointer is a real nil-deref crash.

## Watch out

- Empty slices may yield a nil data pointer.
- Guard `len(s) == 0` before reading element 0.
