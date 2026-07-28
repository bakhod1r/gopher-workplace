# Nil guards apply to unsafe pointers too

## The idea

`unsafe.Pointer` can be nil; dereferencing it panics like any nil pointer, so the guard must come before the read.

## Why it matters

Skipping the guard on unsafe reads is a real nil-deref crash.

## Watch out

- Reorder: nil check, then dereference.
- unsafe doesn't exempt you from nil safety.
