# Cartesian Product

## Intuition

Order is part of the contract here: swapping the loops produces the same pairs in a different sequence, which would silently change test output.

## Approach

1. Allocate with capacity `len(as)*len(bs)`.
2. For each `a`, append a pair with every `b`.

## Solution

```go
func Product[A, B any](as []A, bs []B) []Pair[A, B] {
	out := make([]Pair[A, B], 0, len(as)*len(bs))
	for _, a := range as {
		for _, b := range bs {
			out = append(out, Pair[A, B]{First: a, Second: b})
		}
	}
	return out
}
```

## Walkthrough

`Product([]int{1,2}, []string{"a"})` emits `{1 a}` then `{2 a}` — `as` varies slowest.

## Pitfalls

- Swapping the loops and producing b-major order.
- Growing the slice element by element without the capacity hint.
- Returning nil when one input is empty.
