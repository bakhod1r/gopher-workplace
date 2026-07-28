# Arrays as map keys

## The idea

Unlike slices, **arrays are comparable**, so a `[2]int` (or `[N]T` of comparable
T) is a valid map key — perfect for coordinate pairs:

```go
m[[2]int{c[0], c[1]}]++
```

The array's element order *is* its identity; swapping fields makes a different
key.

## Why it matters

Keying by a fixed-size tuple (grid cells, RGB, small composite keys) avoids
string concatenation and its ambiguities. Getting the field order right is the
whole correctness of the grouping.

## Watch out

- Slices, maps, and funcs are **not** comparable — can't be keys.
- A struct of comparable fields is also a valid key.
- Array keys are compared and hashed by value (all elements).
