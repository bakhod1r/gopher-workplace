// Package speedupratio — Gopher Workplace challenge.
package speedupratio

// Speedup returns how many times faster candidate is than base: a base of
// 100ns against a candidate of 25ns is 4. A non-positive input gives 0.
//
// Examples:
//
//	Speedup(100, 25) => 4
func Speedup(baseNS, candidateNS float64) float64 {
	panic("not implemented")
}

// PercentChange returns the change from base to candidate as a signed
// percentage of base: negative means faster. A non-positive base gives 0.
//
// Examples:
//
//	PercentChange(100, 80) => -20
func PercentChange(baseNS, candidateNS float64) float64 {
	panic("not implemented")
}
