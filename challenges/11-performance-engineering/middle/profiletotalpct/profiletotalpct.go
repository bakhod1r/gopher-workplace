// Package profiletotalpct — Gopher Workplace challenge.
package profiletotalpct

// Row is one line of a pprof top listing, carrying both percentage columns.
type Row struct {
	Func   string
	Flat   int64
	Pct    float64
	CumPct float64
}

// Top renders the listing pprof prints: rows ordered by flat descending then
// name ascending, each carrying its own share of the total (Pct) and the
// running total of the shares so far (CumPct). Both percentages are rounded
// to two decimals, non-positive values are dropped, and an empty profile
// gives an empty, non-nil slice.
//
// Examples:
//
//	Top({"a": 3, "b": 1}) => [{a 3 75 75} {b 1 25 100}]
func Top(flat map[string]int64) []Row {
	panic("not implemented")
}

// CoveringCount returns how many of those rows are needed to account for at
// least pct percent of the profile — the "80% of the time is in N functions"
// number. A pct at or below 0 needs no rows; more than 100 needs them all.
//
// Examples:
//
//	CoveringCount({"a": 3, "b": 1}, 50) => 1
func CoveringCount(flat map[string]int64, pct float64) int {
	panic("not implemented")
}
