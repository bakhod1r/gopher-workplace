# Writing through a pointer

## The idea

`*p = v` stores v at the address p holds, mutating the caller's variable.

## Why it matters

Reset/clear helpers rely on pointer (or slice) parameters.

## Watch out

- Only `*p = ...` reaches the pointee.
- Passing the value by copy would not clear the original.
