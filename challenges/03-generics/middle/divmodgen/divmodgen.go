// Package divmodgen — Gopher Workplace challenge.
package divmodgen

// Integer is the set of signed integer types used here.
type Integer interface {
	~int | ~int64
}

// DivMod returns a/b and a%b with Go's truncating semantics,
// reporting false when b is zero.
func DivMod[T Integer](a, b T) (T, T, bool) {
	// TODO(candidate): guard division by zero, then divide.
	panic("not implemented")
}
