// Package regressionflag — Gopher Workplace challenge.
package regressionflag

// Classify turns a signed percentage change into a verdict, given a
// tolerance band that counts as measurement noise:
//
//	"regression"  change above +tolerance
//	"improvement" change below -tolerance
//	"noise"       anything within the band, boundaries included
//
// A negative tolerance is treated as 0.
//
// Examples:
//
//	Classify(-20, 5) => "improvement"
func Classify(percent, tolerance float64) string {
	panic("not implemented")
}

// Failing reports whether a set of changes contains a regression at the given
// tolerance — the check a CI job runs over a whole benchmark suite.
//
// Examples:
//
//	Failing([]float64{-10, 2}, 5) => false
func Failing(percents []float64, tolerance float64) bool {
	panic("not implemented")
}
