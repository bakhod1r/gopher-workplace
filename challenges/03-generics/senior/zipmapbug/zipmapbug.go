// Package zipmapbug — Gopher Workplace challenge.
package zipmapbug

// ZipMap pairs keys with values positionally, stopping at the
// shorter slice. Later duplicate keys win.
//
// Examples:
//
//	ZipMap([]string{"a"}, []int{1, 2}) => {a:1}
func ZipMap[K comparable, V any](keys []K, vals []V) map[K]V {
	// CHANGE CODE BELOW THIS LINE
	out := make(map[K]V, len(keys))
	for i, k := range keys {
		out[k] = vals[i]
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
