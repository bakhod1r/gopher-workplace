// Package keystringgen — Gopher Workplace challenge.
package keystringgen

// HeaderName is a header key.
type HeaderName string

// Normalize returns m re-keyed with plain lowercase strings.
func Normalize[K ~string, V any](m map[K]V) map[string]V {
	// TODO(candidate): convert and lowercase every key.
	panic("not implemented")
}
