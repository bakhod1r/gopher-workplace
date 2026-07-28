# Sizeof over arrays

## The idea

An array's size is length times element size and is a compile-time constant; measuring an element gives only one slot.

## Why it matters

Correct total sizes matter for buffer allocation and interop.

## Watch out

- `Sizeof(*p)` is the array; `Sizeof(p[0])` is one element.
- Arrays carry their length in the type, so Sizeof knows the total.
