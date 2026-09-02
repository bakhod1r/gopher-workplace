// Package ptrofgen — Gopher Workplace challenge.
package ptrofgen

// Ptr returns a pointer to a copy of v.
//
// Examples:
//
//	*Ptr(7)      => 7
//	*Ptr("go") => "go"
func Ptr[T any](v T) *T {
	// TODO(candidate): return the address of the parameter.
	panic("not implemented")
}
