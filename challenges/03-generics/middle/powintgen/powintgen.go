// Package powintgen — Gopher Workplace challenge.
package powintgen

// Integer is the set of signed integer types used here.
type Integer interface {
	~int | ~int64
}

// Pow returns base raised to exp using repeated squaring.
// It returns 1 for exp == 0 and 0 for a negative exp.
func Pow[T Integer](base T, exp int) T {
	// TODO(candidate): raise base to exp by repeated squaring.
	panic("not implemented")
}
