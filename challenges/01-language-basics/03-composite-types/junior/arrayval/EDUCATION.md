# Arrays are values

## The idea

A fixed-size array (`[N]T`) is a value type: its length is part of its type, and
assigning or passing it **copies** every element. This differs from a slice,
which shares a backing array.

## Why it matters

Passing an array to a function copies it, so mutations inside don't affect the
caller. To mutate, pass a pointer (`*[N]T`) or use a slice.

## Watch out

- `[3]int` and `[4]int` are different types.
- Arrays are comparable with `==` (element-wise); slices are not.
- Large arrays copied by value are expensive.
