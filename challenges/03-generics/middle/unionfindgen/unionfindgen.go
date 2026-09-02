// Package unionfindgen — Gopher Workplace challenge.
package unionfindgen

// Disjoint tracks which elements belong to the same set.
// Use NewDisjoint to create one.
type Disjoint[T comparable] struct {
	parent map[T]T
}

// NewDisjoint returns an empty structure.
func NewDisjoint[T comparable]() *Disjoint[T] {
	// TODO(candidate): allocate the parent map.
	panic("not implemented")
}

// Find returns the representative of v's set, adding v when unseen.
func (d *Disjoint[T]) Find(v T) T {
	// TODO(candidate): walk to the root, compressing the path.
	panic("not implemented")
}

// Union merges the sets containing a and b.
func (d *Disjoint[T]) Union(a, b T) {
	// TODO(candidate): point one root at the other.
	panic("not implemented")
}

// Connected reports whether a and b are in the same set.
func (d *Disjoint[T]) Connected(a, b T) bool {
	// TODO(candidate): compare the two representatives.
	panic("not implemented")
}
