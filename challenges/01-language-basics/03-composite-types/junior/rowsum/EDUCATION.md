# 2-D data as slices of slices

## The idea

Go's 2-D structures are slices of slices, and they can be **ragged**. Nested
ranging handles that:

```go
for _, row := range grid {
	s := 0
	for _, v := range row { s += v }
	out = append(out, s)
}
```

## Why it matters

Grids, matrices, and tables are modeled this way. Rows are independent slices, so
they need not be rectangular.

## Watch out

- Rows can differ in length; don't assume a fixed width.
- An empty row sums to 0.
- Rows are separate allocations and may be shared.
