// Package cumdoublecountbug — Gopher Workplace challenge.
package cumdoublecountbug

// Sample is one profile sample: the call stack, caller first, and its value.
type Sample struct {
	Stack []string
	Value int64
}

// CumSum totals the cumulative time of every function: a sample's value is
// credited to each function on its stack, but a function appearing more than
// once in one stack — recursion — is credited only once for that sample.
//
// Examples:
//
//	CumSum([{["main","a","b"], 5}]) => {"main":5, "a":5, "b":5}
func CumSum(samples []Sample) map[string]int64 {
	out := make(map[string]int64)
	for _, s := range samples {
		if s.Value <= 0 || len(s.Stack) == 0 {
			continue
		}
		seen := make(map[string]bool, len(s.Stack))
		for _, fn := range s.Stack {
			// CHANGE CODE BELOW THIS LINE
			seen[fn] = true
			// CHANGE CODE ABOVE THIS LINE
			out[fn] += s.Value
		}
	}
	return out
}
