// Package alignof reports the alignment of int64 on the platform. A planted bug
// returns the SIZE instead of the ALIGNMENT (they can differ for composite
// types; the test pins the int64 case where both are 8 but the API must be
// Alignof).
package alignof

import "unsafe"

type S struct {
	A bool
	B [3]int32
}

// FieldAlign returns the alignment requirement of field B.
func FieldAlign(s *S) uintptr {
	// CHANGE CODE BELOW THIS LINE
	return unsafe.Sizeof(s.B)
	// CHANGE CODE ABOVE THIS LINE
}
