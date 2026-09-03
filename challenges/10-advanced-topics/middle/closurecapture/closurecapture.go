// Package closurecapture — Gopher Workplace challenge.
package closurecapture

// Counter returns a function that yields start, start+1, start+2 and so
// on, one value per call.
//
// The captured variable outlives Counter's frame, so it must live on the
// heap — that is what a closure over a mutable local costs.
//
// Examples:
//
//	c := Counter(1); c(), c() => 1, 2
func Counter(start int) func() int {
	panic("not implemented")
}
