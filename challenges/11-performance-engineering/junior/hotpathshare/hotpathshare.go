// Package hotpathshare — Gopher Workplace challenge.
package hotpathshare

// Share reports what fraction of the whole profile the named functions
// account for, as a value in [0,1]. Names missing from flat contribute
// nothing, a name listed twice is counted once, and an empty profile is 0.
//
// Examples:
//
//	Share({"a":3,"b":1}, ["a"]) => 0.75
func Share(flat map[string]int64, names []string) float64 {
	panic("not implemented")
}
