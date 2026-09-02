// Package unitmulgen — Gopher Workplace challenge.
package unitmulgen

// Meters is a distance.
type Meters float64

// Seconds is a duration.
type Seconds float64

// Times multiplies a unit value by a plain factor.
func Times[T ~float64](v T, factor float64) T {
	// TODO(candidate): multiply by the factor, keeping the unit.
	panic("not implemented")
}

// SumUnits totals unit values, keeping the unit.
func SumUnits[T ~float64](s []T) T {
	// TODO(candidate): add every element.
	panic("not implemented")
}
