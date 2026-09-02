// Package sumvalues — Gopher Workplace challenge.
package sumvalues

// Number is the set of numeric types this package works with.
type Number interface {
	~int | ~int64 | ~float64
}

// SumValues returns the total of the values in m.
//
// Examples:
//
//	SumValues(map[string]int{"a": 1, "b": 2}) => 3
func SumValues[K comparable, V Number](m map[K]V) V {
	// TODO(candidate): add every value into a running total.
	panic("not implemented")
}
