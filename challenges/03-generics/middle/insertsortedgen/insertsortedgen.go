// Package insertsortedgen — Gopher Workplace challenge.
package insertsortedgen

import (
	"cmp"
)

// InsertSorted returns a copy of the sorted slice s with v inserted
// so the result stays sorted. Equal elements place v last.
func InsertSorted[T cmp.Ordered](s []T, v T) []T {
	// TODO(candidate): find the insertion point, then insert into a copy.
	panic("not implemented")
}
