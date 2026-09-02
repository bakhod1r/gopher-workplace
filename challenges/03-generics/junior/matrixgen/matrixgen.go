// Package matrixgen — Gopher Workplace challenge.
package matrixgen

// Matrix is a rows x cols grid of T stored in one flat slice.
// Use NewMatrix to create one.
type Matrix[T any] struct {
	rows  int
	cols  int
	cells []T
}

// NewMatrix returns a rows x cols matrix of zero values.
func NewMatrix[T any](rows, cols int) *Matrix[T] {
	// TODO(candidate): allocate the backing storage.
	panic("not implemented")
}

// At returns the cell at (r, c) and true, or the zero value and
// false when the position is out of range.
func (m *Matrix[T]) At(r, c int) (T, bool) {
	// TODO(candidate): bounds-check, then read the flat index.
	panic("not implemented")
}

// Set writes v at (r, c) and reports whether the write happened.
func (m *Matrix[T]) Set(r, c int, v T) bool {
	// TODO(candidate): bounds-check, then write the flat index.
	panic("not implemented")
}
