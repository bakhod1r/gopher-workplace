// Package nilslicedata returns the first element's value or a default for an
// empty slice, using unsafe.SliceData. A planted bug dereferences the data
// pointer without checking length; for an empty slice the data pointer may be
// nil, panicking.
package nilslicedata

import "unsafe"

// FirstOr returns s[0] via the data pointer, or def when s is empty.
func FirstOr(s []int, def int) int {
	// CHANGE CODE BELOW THIS LINE
	return *unsafe.SliceData(s)
	// CHANGE CODE ABOVE THIS LINE
}
