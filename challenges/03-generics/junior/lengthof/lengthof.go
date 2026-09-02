// Package lengthof — Gopher Workplace challenge.
package lengthof

// Sized is the set of types len() accepts here.
type Sized interface {
	~string | ~[]byte
}

// Length returns the number of bytes in v.
//
// Examples:
//
//	Length("abc")          => 3
//	Length([]byte{1, 2})   => 2
func Length[T Sized](v T) int {
	// TODO(candidate): return the length of the value.
	panic("not implemented")
}
