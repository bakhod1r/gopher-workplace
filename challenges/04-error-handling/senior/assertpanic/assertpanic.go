// Package assertpanic — Gopher Workplace challenge.
package assertpanic

// Panicked runs f and reports whether it panicked, with the payload.
//
// Examples:
//
//	Panicked(func() {})              => nil, false
//	Panicked(func() { panic("x") })  => "x", true
func Panicked(f func()) (value any, panicked bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
