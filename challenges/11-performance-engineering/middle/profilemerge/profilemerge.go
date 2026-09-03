// Package profilemerge — Gopher Workplace challenge.
package profilemerge

// Profile is one collected profile: per-function values plus the number of
// samples behind them.
type Profile struct {
	Flat    map[string]int64
	Samples int64
}

// Merge combines profiles collected from several machines or several runs
// into one, summing both the per-function values and the sample counts. Nil
// or empty profiles contribute nothing, the inputs must not be modified, and
// merging nothing gives a profile with an empty, non-nil map.
//
// Examples:
//
//	Merge([{Flat:{"a":1}, Samples:1}, {Flat:{"a":2}, Samples:1}]) => {{"a":3}, 2}
func Merge(profiles []Profile) Profile {
	panic("not implemented")
}

// Total returns the sum of every value in a profile.
//
// Examples:
//
//	Total(Profile{Flat: map[string]int64{"a": 1, "b": 2}}) => 3
func Total(p Profile) int64 {
	panic("not implemented")
}
