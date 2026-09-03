// Package profilesort — Gopher Workplace challenge.
package profilesort

// Entry is one row of a profile listing.
type Entry struct {
	Func string
	Flat int64
	Cum  int64
}

// SortByCum returns a new slice ordered by Cum descending, then Flat
// descending, then Func ascending. The input slice must not be modified.
//
// Examples:
//
//	SortByCum([{a 1 5} {b 2 9}]) => [{b 2 9} {a 1 5}]
func SortByCum(entries []Entry) []Entry {
	panic("not implemented")
}
