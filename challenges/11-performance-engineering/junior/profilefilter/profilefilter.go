// Package profilefilter — Gopher Workplace challenge.
package profilefilter

// Entry is one row of a profile listing.
type Entry struct {
	Func  string
	Value int64
}

// Filter drops the rows that account for less than minPct percent of total,
// the job pprof's -nodefraction flag does. Order is preserved, rows at
// exactly the threshold are kept, and a non-positive total keeps nothing.
//
// Examples:
//
//	Filter([{a 50} {b 1}], 100, 5) => [{a 50}]
func Filter(entries []Entry, total int64, minPct float64) []Entry {
	panic("not implemented")
}
