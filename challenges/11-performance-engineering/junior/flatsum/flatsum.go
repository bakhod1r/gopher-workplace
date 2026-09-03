// Package flatsum — Gopher Workplace challenge.
package flatsum

// Sample is one CPU profile sample: the function that was executing and the
// nanoseconds attributed to it.
type Sample struct {
	Func string
	Self int64
}

// FlatSum totals the self time of every function, the "flat" column of
// `go tool pprof -top`. Samples with a non-positive Self are ignored, and a
// function with no remaining samples must not appear in the result.
//
// Examples:
//
//	FlatSum([{"a",3},{"b",1},{"a",2}]) => {"a":5, "b":1}
func FlatSum(samples []Sample) map[string]int64 {
	panic("not implemented")
}
