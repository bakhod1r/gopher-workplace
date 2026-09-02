// Package shallowclonebug — Gopher Workplace challenge.
package shallowclonebug

// Doc is a document with a list of tags.
type Doc[T any] struct {
	Title string
	Tags  []T
}

// CloneDoc returns an independent copy of d.
//
// Mutating the clone's Tags must never be visible through d.
//
// Examples:
//
//	c := CloneDoc(d); c.Tags[0] = "x" // d.Tags[0] unchanged
func CloneDoc[T any](d Doc[T]) Doc[T] {
	// CHANGE CODE BELOW THIS LINE
	out := d
	return out
	// CHANGE CODE ABOVE THIS LINE
}

// CloneAll returns independent copies of every document in ds.
func CloneAll[T any](ds []Doc[T]) []Doc[T] {
	out := make([]Doc[T], len(ds))
	for i, d := range ds {
		out[i] = CloneDoc(d)
	}
	return out
}
