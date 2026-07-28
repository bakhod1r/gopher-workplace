// Package mapmerge merges two maps.
package mapmerge

// Merge returns a new map with all entries of a, then b overriding on key
// collisions. Inputs are not modified.
//
// TODO(candidate): build a fresh map and copy both in.
func Merge(a, b map[string]int) map[string]int {
	panic("not implemented")
}
