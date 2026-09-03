// Package profilepercent — Gopher Workplace challenge.
package profilepercent

// Percent renders a pprof percentage column: value as a share of total,
// rounded to two decimal places. A non-positive total is 0.
//
// Examples:
//
//	Percent(1, 3) => 33.33
func Percent(value, total int64) float64 {
	panic("not implemented")
}

// Format renders the same share the way pprof prints it, e.g. "33.33%".
//
// Examples:
//
//	Format(1, 3) => "33.33%"
func Format(value, total int64) string {
	panic("not implemented")
}
