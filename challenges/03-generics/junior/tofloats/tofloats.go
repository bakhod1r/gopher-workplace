// Package tofloats — Gopher Workplace challenge.
package tofloats

// Integer is the set of integer types.
type Integer interface {
	~int | ~int64
}

// ToFloats converts every element of s to float64.
//
// Examples:
//
//	ToFloats([]int{1, 2}) => []float64{1, 2}
func ToFloats[T Integer](s []T) []float64 {
	// TODO(candidate): convert each element to float64.
	panic("not implemented")
}
