// Package zipmapgen — Gopher Workplace challenge.
package zipmapgen

// ZipMap pairs keys with values positionally, stopping at the
// shorter slice. Later duplicate keys win.
func ZipMap[K comparable, V any](keys []K, vals []V) map[K]V {
	// TODO(candidate): pair matching positions into a map.
	panic("not implemented")
}
