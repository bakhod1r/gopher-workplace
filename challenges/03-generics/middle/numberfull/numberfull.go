// Package numberfull — Gopher Workplace challenge.
package numberfull

// Number covers signed, unsigned and floating-point types.
type Number interface {
	~int | ~int64 | ~uint | ~uint64 | ~float32 | ~float64
}

// Total sums s. It accepts signed, unsigned and floating-point
// types alike.
func Total[T Number](s []T) T {
	// TODO(candidate): add every element into a running total.
	panic("not implemented")
}
