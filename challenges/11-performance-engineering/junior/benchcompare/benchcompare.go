// Package benchcompare — Gopher Workplace challenge.
package benchcompare

// Delta is one row of a benchstat-style comparison.
type Delta struct {
	Name    string
	Base    float64
	New     float64
	Percent float64
}

// Compare pairs two sets of benchmark results by name and reports the signed
// percentage change from base to candidate, negative meaning faster. Only
// names present in both sets appear, rows are ordered by name, and a base of
// zero or less is skipped.
//
// Examples:
//
//	Compare({"A":100}, {"A":80}) => [{A 100 80 -20}]
func Compare(base, candidate map[string]float64) []Delta {
	panic("not implemented")
}
