// Package allocsonce — Gopher Workplace challenge.
package allocsonce

// Index is a lazily built lookup over a fixed word list: the map is
// constructed on the first Lookup and reused by every later one, so a hot
// read path allocates nothing after the first call.
type Index struct {
	Words  []string
	byWord map[string]int
}

// Lookup returns the position of w in the index and whether it was found.
// The first call builds the map; later calls must not allocate.
//
// Examples:
//
//	(&Index{Words: []string{"a", "b"}}).Lookup("b") => 1, true
func (ix *Index) Lookup(w string) (int, bool) {
	panic("not implemented")
}
