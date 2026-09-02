// Package typeswitchgen — Gopher Workplace challenge.
package typeswitchgen

// Describe reports what kind of value v is.
// A type switch needs an interface value, so v is converted first.
func Describe[T any](v T) string {
	// TODO(candidate): convert to any, then type-switch.
	panic("not implemented")
}
