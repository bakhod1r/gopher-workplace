// Package selftimeleakbug — Gopher Workplace challenge.
package selftimeleakbug

// Sample is one profile sample: the call stack, caller first, and its value.
type Sample struct {
	Stack []string
	Value int64
}

// SelfTime totals the self time of every function: only the frame that was
// actually executing — the leaf of the stack — is credited.
//
// Examples:
//
//	SelfTime([{["main","a","b"], 5}]) => {"b":5}
func SelfTime(samples []Sample) map[string]int64 {
	out := make(map[string]int64)
	for _, s := range samples {
		if s.Value <= 0 || len(s.Stack) == 0 {
			continue
		}
		// CHANGE CODE BELOW THIS LINE
		frame := s.Stack[0]
		// CHANGE CODE ABOVE THIS LINE
		out[frame] += s.Value
	}
	return out
}
