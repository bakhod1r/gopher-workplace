# Testing pointers for nil

## The idea

The zero value of any pointer is nil; comparing with `== nil` detects unset references before use.

## Why it matters

Nil-counting/validation guards precede bulk dereferencing.

## Watch out

- Only `== nil` distinguishes an unset pointer.
- Dereferencing the nil ones would panic.
