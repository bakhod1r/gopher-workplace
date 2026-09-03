// Package callgraphedge — Gopher Workplace challenge.
package callgraphedge

// Sample is one profile sample: the call stack, caller first, and its value.
type Sample struct {
	Stack []string
	Value int64
}

// Edge is one caller-to-callee arc of the call graph, carrying the value that
// flowed along it.
type Edge struct {
	Caller string
	Callee string
	Value  int64
}

// Edges builds the weighted call graph pprof draws: one arc per adjacent
// pair of frames, summed across samples. Repeated adjacent pairs inside one
// stack (direct recursion) are counted once per occurrence. Arcs are ordered
// by value descending, then caller ascending, then callee ascending. Samples
// with a non-positive value or fewer than two frames contribute nothing.
//
// Examples:
//
//	Edges([{["a","b","c"], 5}]) => [{a b 5} {b c 5}]
func Edges(samples []Sample) []Edge {
	panic("not implemented")
}

// CalleesOf returns the functions called directly by caller, ordered by the
// value on the arc descending, then by name ascending.
//
// Examples:
//
//	CalleesOf([{a b 5} {a c 9}], "a") => ["c", "b"]
func CalleesOf(edges []Edge, caller string) []string {
	panic("not implemented")
}
