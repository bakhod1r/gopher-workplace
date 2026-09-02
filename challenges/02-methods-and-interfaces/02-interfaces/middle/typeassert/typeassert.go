// Package typeassert — Gopher Workplace challenge.
package typeassert

// AsInt extracts an int payload.
//
// Examples:
//
//	AsInt(5)   => 5, true
//	AsInt("5") => 0, false
func AsInt(v any) (int, bool) {
	// TODO(candidate): comma-ok assertion.
	panic("not implemented")
}

// SumInts totals every int in the slice, ignoring other types.
//
// Examples:
//
//	SumInts([]any{1, "x", 2}) => 3
func SumInts(vs []any) int {
	// TODO(candidate): add up only the ints.
	panic("not implemented")
}
