# Matrix multiplication

## Intuition

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

## Approach

1. m=len(a); if 0 return nil.
2. n=len(a[0]); if len(b)!=n return nil.
3. p=len(b[0]).
4. For each i,j sum over k of a[i][k]*b[k][j].
5. Return m x p result.

## Solution

```go
func Mul(a, b [][]int) [][]int {
	m := len(a)
	if m == 0 {
		return nil
	}
	n := len(a[0])
	if len(b) != n {
		return nil
	}
	p := 0
	if n > 0 {
		p = len(b[0])
	}
	out := make([][]int, m)
	for i := range out {
		out[i] = make([]int, p)
		for j := 0; j < p; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				sum += a[i][k] * b[k][j]
			}
			out[i][j] = sum
		}
	}
	return out
}
```

## Walkthrough

out[0][0]=1*5+2*7=19; out[0][1]=1*6+2*8=22; out[1][0]=3*5+4*7=43; out[1][1]=3*6+4*8=50.

## Pitfalls

- Validate `len(a[0]) == len(b)` before multiplying.
- Pre-allocate the m×p result.
- Naive triple loop is O(m·n·p); fine for small matrices.
