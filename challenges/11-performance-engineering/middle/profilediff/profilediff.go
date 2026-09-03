// Package profilediff — Gopher Workplace challenge.
package profilediff

// Change is one row of a profile-to-profile diff.
type Change struct {
	Func  string
	Base  int64
	New   int64
	Delta int64
}

// Diff compares two flat profiles over the union of their function names: a
// function missing from one side counts as zero there, because "it stopped
// appearing" is exactly the signal you are looking for. Rows with a zero
// delta are dropped, and the result is ordered by absolute delta descending,
// then by name ascending.
//
// Examples:
//
//	Diff({"a":10}, {"a":4}) => [{a 10 4 -6}]
func Diff(base, candidate map[string]int64) []Change {
	panic("not implemented")
}
