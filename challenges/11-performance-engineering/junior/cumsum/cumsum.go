// Package cumsum — Gopher Workplace challenge.
package cumsum

// Sample is one profile sample: the call stack that was on the CPU, caller
// first, plus the nanoseconds it accounts for.
type Sample struct {
	Stack []string
	Value int64
}

// CumSum totals the cumulative time of every function — the "cum" column of
// `go tool pprof -top`. A sample's value is credited to every function on its
// stack, but a function that appears more than once in one stack (recursion)
// is credited only once for that sample. Samples with a non-positive Value or
// an empty stack are ignored.
//
// Examples:
//
//	CumSum([{["main","a","b"], 5}]) => {"main":5, "a":5, "b":5}
func CumSum(samples []Sample) map[string]int64 {
	panic("not implemented")
}
