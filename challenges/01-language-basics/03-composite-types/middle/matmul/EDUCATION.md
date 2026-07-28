# Matrix multiplication

## The idea

Each output cell is a dot product of a row of `a` and a column of `b`:

```go
for i := range a {
	for j := range b[0] {
		s := 0
		for k := range b { s += a[i][k] * b[k][j] }
		out[i][j] = s
	}
}
```

## Why it matters

Beyond math, it's the archetype of nested iteration over 2-D data and dimension
bookkeeping. The inner dimension of `a` must equal the outer of `b`.

## Watch out

- Validate `len(a[0]) == len(b)` before multiplying.
- Pre-allocate the m×p result.
- Naive triple loop is O(m·n·p); fine for small matrices.
