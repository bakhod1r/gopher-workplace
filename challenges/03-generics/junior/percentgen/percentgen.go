// Package percentgen — Gopher Workplace challenge.
package percentgen

// Number is the set of numeric types this package works with.
type Number interface {
	~int | ~int64 | ~float64
}

// Percent returns part as a percentage of whole.
// It returns 0 when whole is zero.
//
// Examples:
//
//	Percent(1, 4)  => 25
//	Percent(1, 0)  => 0
func Percent[T Number](part, whole T) float64 {
	// TODO(candidate): convert to float64 and scale by 100.
	panic("not implemented")
}
