// Package tildeint — Gopher Workplace challenge.
package tildeint

// IntLike accepts any type whose underlying type is int.
type IntLike interface {
	~int
}

// Celsius is a temperature in degrees Celsius.
type Celsius int

// SumTemps totals a slice of temperature-like values.
//
// Examples:
//
//	SumTemps([]Celsius{1, 2}) => Celsius(3)
func SumTemps[T IntLike](s []T) T {
	// TODO(candidate): add every element into a running total.
	panic("not implemented")
}
