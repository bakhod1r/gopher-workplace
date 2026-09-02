// Package comparablekeygen — Gopher Workplace challenge.
package comparablekeygen

// Index builds a lookup keyed by a comparable type.
// The constraint makes an uncomparable key a compile error
// rather than a run-time panic.
func Index[K comparable, V any](keys []K, vals []V) map[K]V {
	// TODO(candidate): pair the slices into a map.
	panic("not implemented")
}

// IndexAny does the same with interface keys. Storing an
// uncomparable value panics at run time. It is provided for
// comparison.
func IndexAny(keys []any, vals []any) map[any]any {
	n := len(keys)
	if len(vals) < n {
		n = len(vals)
	}
	out := make(map[any]any, n)
	for i := 0; i < n; i++ {
		out[keys[i]] = vals[i]
	}
	return out
}
