// Package enumerifc — Gopher Workplace challenge.
package enumerifc

// Enumerable exposes items by index.
type Enumerable interface {
	Len() int
	At(i int) string
}

// Words is a list of words.
type Words []string

// Len returns the number of words.
//
// Examples:
//
//	Words{"a", "b"}.Len() => 2
func (w Words) Len() int {
	// TODO(candidate): return the length.
	panic("not implemented")
}

// At returns the word at index i.
//
// Examples:
//
//	Words{"a", "b"}.At(1) => "b"
func (w Words) At(i int) string {
	// TODO(candidate): return the element.
	panic("not implemented")
}

// Join concatenates every item with sep between them.
func Join(e Enumerable, sep string) string {
	// TODO(candidate): walk indexes 0..Len()-1.
	panic("not implemented")
}
