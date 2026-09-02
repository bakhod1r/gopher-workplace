// Package itabcost — Gopher Workplace challenge.
package itabcost

// Op folds one value into an accumulator.
type Op interface {
	Apply(acc, v int) int
}

// AddOp adds the value plus N.
type AddOp struct {
	N int
}

// Apply returns acc + v + N.
//
// Examples:
//
//	AddOp{N: 2}.Apply(0, 1) => 3
func (o AddOp) Apply(acc, v int) int {
	// TODO(candidate): acc + v + N.
	panic("not implemented")
}

// MulOp multiplies the accumulator by the value.
type MulOp struct{}

// Apply returns acc * v.
func (MulOp) Apply(acc, v int) int {
	// TODO(candidate): acc * v.
	panic("not implemented")
}

// RunIface folds vs through op, starting from start.
func RunIface(op Op, start int, vs []int) int {
	// TODO(candidate): fold through the interface.
	panic("not implemented")
}

// RunConcrete folds vs through a concrete AddOp, with no dynamic dispatch.
func RunConcrete(op AddOp, start int, vs []int) int {
	// TODO(candidate): fold through the concrete type.
	panic("not implemented")
}
