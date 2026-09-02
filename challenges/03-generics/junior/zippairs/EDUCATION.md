# Zip To Pairs

## Intuition

Once instantiated, `Pair[A, B]` behaves like any struct type: it can be a slice element, a map value, or a field.

## Approach

1. Take the smaller of the two lengths.
2. Allocate `[]Pair[A, B]` with that capacity.
3. Append a pair per position.

## Solution

```go
func Zip[A, B any](a []A, b []B) []Pair[A, B] {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	out := make([]Pair[A, B], 0, n)
	for i := 0; i < n; i++ {
		out = append(out, Pair[A, B]{First: a[i], Second: b[i]})
	}
	return out
}
```

## Walkthrough

`Zip([]int{1, 2, 3}, []string{"a"})` stops after one pair because `b` runs out first.

## Pitfalls

- Ranging over `a` and indexing `b`, which panics when `b` is shorter.
- Writing `make([]Pair, 0, n)` without the type arguments.
- Padding the shorter slice with zero values.
