// Package structasarray reads both fields of a two-int32 struct by viewing it as
// a [2]int32 array through unsafe.Pointer. A planted bug reads only the first
// element, ignoring the reinterpreted view's second slot.
package structasarray

import "unsafe"

type Pair struct {
	A int32
	B int32
}

// Sum returns A+B by reinterpreting the struct as a [2]int32 array.
func Sum(p *Pair) int32 {
	arr := (*[2]int32)(unsafe.Pointer(p))
	// CHANGE CODE BELOW THIS LINE
	return arr[0]
	// CHANGE CODE ABOVE THIS LINE
}
