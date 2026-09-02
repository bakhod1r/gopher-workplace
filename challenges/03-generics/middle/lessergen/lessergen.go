// Package lessergen — Gopher Workplace challenge.
package lessergen

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

// MaxOf returns the largest element and true, using each element's
// own Less method.
func MaxOf[T Lesser[T]](s []T) (T, bool) {
	// TODO(candidate): track the largest element using Less.
	panic("not implemented")
}
