// Package stackfold — Gopher Workplace challenge.
package stackfold

// Sample is one profile sample: the call stack, caller first, and its value.
type Sample struct {
	Stack []string
	Value int64
}

// Fold renders the samples in the collapsed-stack format flame graph tools
// eat: one line per distinct stack, "frame;frame;frame value". Identical
// stacks are summed, lines are sorted by value descending and then by stack
// ascending, and samples with a non-positive value or an empty stack are
// dropped.
//
// Examples:
//
//	Fold([{["a","b"], 3}, {["a","b"], 2}]) => ["a;b 5"]
func Fold(samples []Sample) []string {
	panic("not implemented")
}
