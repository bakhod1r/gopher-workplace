// Package mapdeepclonegen — Gopher Workplace challenge.
package mapdeepclonegen

// DeepClone copies a map whose values are slices, so the copy
// shares nothing with the original.
func DeepClone[K comparable, V any](m map[K][]V) map[K][]V {
	// TODO(candidate): copy each value slice, not just the map.
	panic("not implemented")
}
