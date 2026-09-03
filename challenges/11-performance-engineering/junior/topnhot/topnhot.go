// Package topnhot — Gopher Workplace challenge.
package topnhot

// Entry is one row of a pprof top listing.
type Entry struct {
	Func  string
	Value int64
}

// TopN returns the n hottest functions in flat, ordered by value descending
// and, for equal values, by name ascending. Fewer than n functions returns
// them all; a non-positive n returns an empty, non-nil slice.
//
// Examples:
//
//	TopN({"a":3,"b":9,"c":3}, 2) => [{b 9} {a 3}]
func TopN(flat map[string]int64, n int) []Entry {
	panic("not implemented")
}
