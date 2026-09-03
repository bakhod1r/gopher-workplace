// Package inlinemark — Gopher Workplace challenge.
package inlinemark

// Frame is one stack frame. Inlined marks a frame the compiler folded into
// its caller: it has no real activation record, and the profiler reconstructs
// it from the inline tree.
type Frame struct {
	Func    string
	Inlined bool
}

// Physical drops the inlined frames, leaving the stack the machine actually
// had — what `pprof -noinlines` shows. The input is not modified, and a stack
// with nothing left gives an empty, non-nil result.
//
// Examples:
//
//	Physical([{a false} {b true} {c false}]) => ["a", "c"]
func Physical(stack []Frame) []string {
	panic("not implemented")
}

// Attribute credits value to the function that owns the leaf's machine code:
// the innermost frame that was not inlined. An all-inlined or empty stack
// credits nothing and returns "", false.
//
// Examples:
//
//	Attribute([{a false} {b true}]) => "a", true
func Attribute(stack []Frame) (string, bool) {
	panic("not implemented")
}
