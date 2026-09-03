// Package contentionrank — Gopher Workplace challenge.
package contentionrank

// Record is one entry of a mutex or block profile: a call site, how many
// times a goroutine blocked there, and the total nanoseconds spent waiting.
type Record struct {
	Site  string
	Count int64
	Delay int64
}

// Site is one ranked row of the report.
type Site struct {
	Site      string
	Count     int64
	Delay     int64
	MeanDelay float64
}

// Rank aggregates the records per call site and orders them by total delay
// descending, then by site ascending. MeanDelay is the delay per blocking
// event. Records with a non-positive count or a negative delay are dropped,
// and sites left with no events do not appear.
//
// Examples:
//
//	Rank([{a 2 10}, {a 2 30}]) => [{a 4 40 10}]
func Rank(records []Record) []Site {
	panic("not implemented")
}

// Worst returns the site with the highest total delay, and false when there
// is none — the one line to act on after reading a mutex profile.
//
// Examples:
//
//	Worst([{a 1 5}, {b 1 50}]) => Site{b 1 50 50}, true
func Worst(records []Record) (Site, bool) {
	panic("not implemented")
}
