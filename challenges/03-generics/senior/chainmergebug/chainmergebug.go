// Package chainmergebug — Gopher Workplace challenge.
package chainmergebug

// Merge combines maps left to right; later maps win on a key clash.
//
// Examples:
//
//	Merge(map[string]int{"a": 1}, map[string]int{"a": 2}) => map[a:2]
func Merge[K comparable, V any](ms ...map[K]V) map[K]V {
	// CHANGE CODE BELOW THIS LINE
	out := make(map[K]V)
	for _, m := range ms {
		for k, v := range m {
			if _, ok := out[k]; !ok {
				out[k] = v
			}
		}
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
