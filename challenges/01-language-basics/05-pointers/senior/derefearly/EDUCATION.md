# Nil checks precede dereferences

## The idea

A dereference evaluates immediately; guarding after it is too late. The nil test must run before touching the pointee.

## Why it matters

Ordering the guard after the access is a real nil-panic bug.

## Watch out

- `head.Val` panics on nil regardless of a later check.
- Put `if head == nil` first.
