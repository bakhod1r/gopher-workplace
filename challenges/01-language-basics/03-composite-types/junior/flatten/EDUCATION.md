# Row-major flattening

## The idea

Flattening a slice of slices concatenates the rows in order:

```go
out := []int{}
for _, row := range grid { out = append(out, row...) }
```

## Why it matters

Converting between nested and flat representations is common (image buffers,
serialization). Row-major order is the default for Go's nested slices.

## Watch out

- `append(out, row...)` spreads the row; without `...` it's a type error.
- Empty and nil rows contribute nothing.
- A single `[]int` indexed by `r*w+c` is more cache-friendly for fixed width.
