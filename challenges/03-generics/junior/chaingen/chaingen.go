// Package chaingen — Gopher Workplace challenge.
package chaingen

// Lookup returns the first value found for k across the maps,
// searched in order, and whether any map had the key.
func Lookup[K comparable, V any](k K, maps ...map[K]V) (V, bool) {
	// TODO(candidate): search the maps in order and return the first hit.
	panic("not implemented")
}
