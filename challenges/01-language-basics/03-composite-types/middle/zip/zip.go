// Package zip pairs two slices into a map.
package zip

// Zip builds a map from keys[i] to vals[i] for the overlapping prefix (min
// length). Extra elements of the longer slice are ignored.
//
// TODO(candidate): iterate to the shorter length.
func Zip(keys []string, vals []int) map[string]int {
	panic("not implemented")
}
