// Package scaleintgen — Gopher Workplace challenge.
package scaleintgen

// Integer is the set of signed integer types used here.
type Integer interface {
	~int | ~int64
}

// Scale maps every element proportionally so the largest becomes
// top. It returns the elements unchanged when the largest is 0.
func Scale[T Integer](s []T, top T) []T {
	// TODO(candidate): divide each element by the largest, scaled to top.
	panic("not implemented")
}
