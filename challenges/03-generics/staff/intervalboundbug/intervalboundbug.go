// Package intervalboundbug — Gopher Workplace challenge.
package intervalboundbug

// Number is any type Accum can sum.
type Number interface {
	~int | ~int64 | ~float64
}

// Accum accumulates values over half-open integer ranges.
type Accum[T Number] struct {
	m map[int]T
}

// Add adds v to every integer position in the half-open range [lo, hi).
// An empty or reversed range adds nothing.
//
// Examples:
//
//	Add(0, 3, 1) touches 0, 1 and 2 — not 3
func (a *Accum[T]) Add(lo, hi int, v T) {
	// CHANGE CODE BELOW THIS LINE
	if a.m == nil {
		a.m = make(map[int]T)
	}
	for i := lo; i <= hi; i++ {
		a.m[i] += v
	}
	// CHANGE CODE ABOVE THIS LINE
}

// At returns the accumulated value at position x.
func (a *Accum[T]) At(x int) T {
	return a.m[x]
}

// Total returns the sum of every accumulated position.
func (a *Accum[T]) Total() T {
	var sum T
	for _, v := range a.m {
		sum += v
	}
	return sum
}

// Touched reports how many positions hold a value.
func (a *Accum[T]) Touched() int {
	return len(a.m)
}
