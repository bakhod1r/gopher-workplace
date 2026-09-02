// Package ptrmethodsetbug — Gopher Workplace challenge.
package ptrmethodsetbug

// Resettable constrains PT to be a pointer to T that can reset itself.
type Resettable[T any] interface {
	*T
	Reset()
}

// ResetAll resets every element of s in place.
// PT constrains *T to carry the Reset method, so T itself need not.
//
// Examples:
//
//	ResetAll[counter, *counter](s) // every s[i].n becomes 0
func ResetAll[T any, PT Resettable[T]](s []T) {
	// CHANGE CODE BELOW THIS LINE
	for _, v := range s {
		PT(&v).Reset()
	}
	// CHANGE CODE ABOVE THIS LINE
}
