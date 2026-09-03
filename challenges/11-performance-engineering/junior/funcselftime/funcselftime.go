// Package funcselftime — Gopher Workplace challenge.
package funcselftime

// Sample is one profile sample: the call stack, caller first, and the value
// it accounts for.
type Sample struct {
	Stack []string
	Value int64
}

// SelfTime totals the self time of every function: only the leaf of each
// stack — the frame that was actually executing — is credited. Samples with a
// non-positive Value or an empty stack are ignored.
//
// Examples:
//
//	SelfTime([{["main","a","b"], 5}]) => {"b":5}
func SelfTime(samples []Sample) map[string]int64 {
	panic("not implemented")
}
