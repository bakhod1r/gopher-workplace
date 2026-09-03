// Package profilenormalize — Gopher Workplace challenge.
package profilenormalize

// Rate converts a profile's raw values into per-second rates, so profiles
// collected over different durations can be compared. A non-positive
// duration gives an empty, non-nil map, and the input is not modified.
//
// Examples:
//
//	Rate({"a": 60}, 30) => {"a": 2}
func Rate(flat map[string]int64, seconds float64) map[string]float64 {
	panic("not implemented")
}

// Fractions converts the values into shares of the profile total, each in
// [0,1] and summing to 1. Non-positive values are dropped, and a total of
// zero gives an empty, non-nil map.
//
// Examples:
//
//	Fractions({"a": 3, "b": 1}) => {"a": 0.75, "b": 0.25}
func Fractions(flat map[string]int64) map[string]float64 {
	panic("not implemented")
}
