// Package rangecopy — Gopher Workplace challenge.
package rangecopy

// Counter is one element of the slice.
type Counter struct {
	N   int
	Pad [64]byte
}

// Bump increments every counter in items, in place.
//
// Ranging by value copies each element; the increment has to reach the
// slice's own storage.
//
// Examples:
//
//	items := []Counter{{N: 1}}; Bump(items) => items[0].N == 2
func Bump(items []Counter) {
	// CHANGE CODE BELOW THIS LINE
	for _, c := range items {
		c.N++
	}
	// CHANGE CODE ABOVE THIS LINE
}
