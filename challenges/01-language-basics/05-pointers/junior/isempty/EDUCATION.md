# nil as the empty case

## The idea

A nil pointer conventionally represents an empty linked structure; testing `== nil` needs no dereference.

## Why it matters

Empty-collection checks precede traversal to avoid nil derefs.

## Watch out

- Comparing to nil is always safe; dereferencing nil is not.
- Many recursive list algorithms base-case on nil.
