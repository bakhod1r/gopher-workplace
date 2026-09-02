// Package accgen — Gopher Workplace challenge.
package accgen

// Number is the set of numeric types this package works with.
type Number interface {
	~int | ~int64 | ~float64
}

// Acc accumulates a running total and count.
// Its zero value is ready to use.
type Acc[T Number] struct {
	total T
	n     int
}

// Add records v.
func (a *Acc[T]) Add(v T) {
	// TODO(candidate): add v to the running total and bump the count.
	panic("not implemented")
}

// Sum returns the running total.
func (a *Acc[T]) Sum() T {
	// TODO(candidate): return the running total.
	panic("not implemented")
}

// Mean returns the average as a float64, or 0 before anything is added.
func (a *Acc[T]) Mean() float64 {
	// TODO(candidate): divide the total by the count, as float64.
	panic("not implemented")
}
