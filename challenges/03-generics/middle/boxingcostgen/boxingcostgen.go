// Package boxingcostgen — Gopher Workplace challenge.
package boxingcostgen

// Number is the numeric set used here.
type Number interface {
	~int | ~int64 | ~float64
}

// SumTyped adds a typed slice with no boxing.
func SumTyped[T Number](s []T) T {
	// TODO(candidate): add every element.
	panic("not implemented")
}

// SumAny adds a slice of interface values, reporting false when
// an element is not an int.
func SumAny(s []any) (int, bool) {
	// TODO(candidate): assert each element, failing on the first non-int.
	panic("not implemented")
}
