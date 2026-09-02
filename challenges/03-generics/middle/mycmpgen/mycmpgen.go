// Package mycmpgen — Gopher Workplace challenge.
package mycmpgen

// Ordered is a hand-written equivalent of cmp.Ordered.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

// Largest returns the largest element and true, using a
// hand-written ordering constraint.
func Largest[T Ordered](s []T) (T, bool) {
	// TODO(candidate): track the largest element.
	panic("not implemented")
}
