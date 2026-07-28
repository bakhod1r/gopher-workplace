// Package shallowclone clones a struct that holds a slice. A planted bug copies
// the struct by value, which shares the slice header (backing array), so editing
// the clone's slice corrupts the original.
package shallowclone

type Doc struct {
	Tags []string
}

// Clone returns an independent copy of d: editing the clone's Tags must not
// affect the original.
func Clone(d *Doc) *Doc {
	// CHANGE CODE BELOW THIS LINE
	cp := *d
	return &cp
	// CHANGE CODE ABOVE THIS LINE
}
