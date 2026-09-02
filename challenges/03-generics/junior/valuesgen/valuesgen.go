// Package valuesgen — Gopher Workplace challenge.
package valuesgen

// Values returns the values of m in unspecified order.
//
// Examples:
//
//	Values(map[string]int{"a": 1}) => []int{1}
//	Values(map[string]int{})        => []int{}
func Values[K comparable, V any](m map[K]V) []V {
	// TODO(candidate): collect every value into a slice.
	panic("not implemented")
}
