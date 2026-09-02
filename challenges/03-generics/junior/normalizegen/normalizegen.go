// Package normalizegen — Gopher Workplace challenge.
package normalizegen

// Float is the set of floating-point types.
type Float interface {
	~float32 | ~float64
}

// Normalize scales s so the largest magnitude becomes 1.
// It returns an empty slice for empty input, and copies s unchanged
// when every element is zero.
//
// Examples:
//
//	Normalize([]float64{2, 4}) => []float64{0.5, 1}
func Normalize[T Float](s []T) []T {
	// TODO(candidate): divide every element by the largest magnitude.
	panic("not implemented")
}
