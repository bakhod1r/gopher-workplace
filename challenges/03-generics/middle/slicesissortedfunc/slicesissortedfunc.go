// Package slicesissortedfunc — Gopher Workplace challenge.
package slicesissortedfunc

// Item is a catalogue entry.
type Item struct {
	Name  string
	Price int
}

// ByName reports whether items are sorted by name.
func ByName(items []Item) bool {
	// TODO(candidate): check the ordering with the stdlib.
	panic("not implemented")
}
