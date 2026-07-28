// Package deepptrfield clones a struct holding a *int. A planted bug copies the
// struct by value, sharing the SAME *int, so editing the clone's pointee changes
// the original.
package deepptrfield

type Box struct {
	P *int
}

// Clone returns an independent copy: writing through the clone's P must not
// affect the original.
func Clone(b *Box) *Box {
	// CHANGE CODE BELOW THIS LINE
	cp := *b
	return &cp
	// CHANGE CODE ABOVE THIS LINE
}
