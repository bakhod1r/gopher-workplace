// Package selfvscum — Gopher Workplace challenge.
package selfvscum

// Sample is one profile sample: the call stack, caller first, and its value.
type Sample struct {
	Stack []string
	Value int64
}

// Node is one function's two profile columns.
type Node struct {
	Flat int64
	Cum  int64
}

// Analyze computes both columns in a single pass: Flat credits only the leaf
// of each stack, Cum credits every distinct frame on it. Samples with a
// non-positive value or an empty stack are ignored.
//
// Examples:
//
//	Analyze([{["main","a"], 5}]) => {"main":{0,5}, "a":{5,5}}
func Analyze(samples []Sample) map[string]Node {
	panic("not implemented")
}

// Leaves returns the functions whose flat time is at least minFlat, ordered
// by flat descending then name ascending — the functions actually burning CPU.
//
// Examples:
//
//	Leaves({"a":{5,5}, "b":{0,5}}, 1) => ["a"]
func Leaves(nodes map[string]Node, minFlat int64) []string {
	panic("not implemented")
}
