# What unsafe.Sizeof measures

## The idea

`Sizeof` returns the size of its operand's type; a pointer is always word-sized regardless of the pointee. Measure the element expression instead.

## Why it matters

Confusing pointer size with pointee size miscomputes buffer strides and layouts.

## Watch out

- `Sizeof(p)` is 8 (a pointer); `Sizeof(p[0])` is the element size.
- Sizeof is a compile-time constant of the type.
