// Package structconv reinterprets a Point as a Vec (identical layout) via
// unsafe. A planted bug builds a Vec from only the first field, dropping the
// second.
package structconv

import "unsafe"

type Point struct{ X, Y int32 }
type Vec struct{ X, Y int32 }

// ToVec reinterprets p as a *Vec (same layout) and returns its value.
func ToVec(p *Point) Vec {
	// CHANGE CODE BELOW THIS LINE
	v := *(*Vec)(unsafe.Pointer(p))
	return Vec{X: v.X}
	// CHANGE CODE ABOVE THIS LINE
}
