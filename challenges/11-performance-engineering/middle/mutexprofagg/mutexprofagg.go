// Package mutexprofagg — Gopher Workplace challenge.
package mutexprofagg

// Contention is one sampled mutex-profile record.
type Contention struct {
	Site  string
	Count int64
	Delay int64
}

// Scale converts sampled counts and delays back to whole-program numbers.
// runtime.SetMutexProfileFraction(n) records one contention event in n, so
// each sample stands for n of them. A fraction at or below 1 records
// everything; a fraction of 0 or less means the profile is off, which is a
// different thing from "no contention" — that case reports ok = false.
//
// Examples:
//
//	Scale(3, 300, 5) => 15, 1500, true
func Scale(count, delay int64, fraction int) (int64, int64, bool) {
	panic("not implemented")
}

// Estimate aggregates the records per site with the sampling correction
// applied, returning the estimated total delay per site. A fraction of 0 or
// less means the profile was off, so nothing can be estimated: the result is
// an empty, non-nil map.
//
// Examples:
//
//	Estimate([{a 1 100}], 5) => {"a": 500}
func Estimate(records []Contention, fraction int) map[string]int64 {
	panic("not implemented")
}

// Confidence reports how trustworthy a site's estimate is, as the number of
// raw samples behind it: fewer than 10 is "low", fewer than 100 is "medium",
// otherwise "high".
//
// Examples:
//
//	Confidence(5) => "low"
func Confidence(samples int64) string {
	panic("not implemented")
}
