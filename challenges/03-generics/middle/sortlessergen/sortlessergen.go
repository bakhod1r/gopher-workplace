// Package sortlessergen — Gopher Workplace challenge.
package sortlessergen

// Lesser is a type that can compare itself with another T.
type Lesser[T any] interface {
	Less(T) bool
}

// Version is an ordered release number.
type Version struct {
	N int
}

// Less reports whether v precedes other.
func (v Version) Less(other Version) bool { return v.N < other.N }

// SortedLess returns a copy of s sorted by each element's Less.
// Equal elements keep their input order.
func SortedLess[T Lesser[T]](s []T) []T {
	// TODO(candidate): clone, then sort stably using Less.
	panic("not implemented")
}
