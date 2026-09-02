// Package rangegen — Gopher Workplace challenge.
package rangegen

// Integer is the set of integer types.
type Integer interface {
	~int | ~int64
}

// Range returns 0, 1, ... n-1 as a slice.
// It returns an empty slice when n <= 0.
//
// Examples:
//
//	Range(3) => []int{0, 1, 2}
//	Range(0) => []int{}
func Range[T Integer](n T) []T {
	// TODO(candidate): build the sequence 0 .. n-1.
	panic("not implemented")
}
