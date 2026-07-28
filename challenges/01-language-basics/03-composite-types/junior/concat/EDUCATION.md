# Variadic parameters and append spread

## The idea

A variadic `...[]int` parameter arrives as a `[][]int`. Flatten with the spread
operator:

```go
out := []int{}
for _, s := range slices { out = append(out, s...) }
```

## Why it matters

Variadic functions give flexible call sites, and `append(out, s...)` is the idiom
for merging slices (`slices.Concat` generalizes it).

## Watch out

- `append(out, s...)` spreads `s`; without `...` it's a type error.
- Appending a `nil` slice is a safe no-op.
- Pre-size with the total length to avoid regrowth.
