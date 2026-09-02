// Package mustswallowbug — Gopher Workplace challenge.
package mustswallowbug

// Must returns v, or panics when err is non-nil.
// It never returns a value produced alongside an error.
//
// Examples:
//
//	cfg := Must(Load(path))
func Must[T any](v T, err error) T {
	// CHANGE CODE BELOW THIS LINE
	if err != nil {
		var zero T
		return zero
	}
	return v
	// CHANGE CODE ABOVE THIS LINE
}
