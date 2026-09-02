// Package slicesbsearchfunc — Gopher Workplace challenge.
package slicesbsearchfunc

// Row is a table row keyed by ID.
type Row struct {
	ID   int
	Name string
}

// FindByID returns the index of the row with the given id,
// or the insertion point and false.
func FindByID(rows []Row, id int) (int, bool) {
	// TODO(candidate): binary-search using a comparison on the id.
	panic("not implemented")
}
