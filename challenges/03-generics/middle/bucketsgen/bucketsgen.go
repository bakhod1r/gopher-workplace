// Package bucketsgen — Gopher Workplace challenge.
package bucketsgen

import (
	"cmp"
)

// Buckets counts how many elements fall into each bucket defined
// by the ascending edges. The result has len(edges)+1 counts.
func Buckets[T cmp.Ordered](s []T, edges []T) []int {
	// TODO(candidate): count the elements below each edge, and above the last.
	panic("not implemented")
}
