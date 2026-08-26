// Package gcmark — Gopher Workplace challenge.
package gcmark

type Object struct {
	Marked bool
	Refs   []*Object
}

func (o *Object) Mark() {
	// TODO(candidate): set Marked to true, recursively mark Refs
	panic("not implemented")
}
