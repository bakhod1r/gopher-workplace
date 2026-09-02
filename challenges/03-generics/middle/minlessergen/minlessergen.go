// Package minlessergen — Gopher Workplace challenge.
package minlessergen

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

// MinOf returns the smallest element and true, using Less.
func MinOf[T Lesser[T]](s []T) (T, bool) {
	// TODO(candidate): track the smallest element using Less.
	panic("not implemented")
}
