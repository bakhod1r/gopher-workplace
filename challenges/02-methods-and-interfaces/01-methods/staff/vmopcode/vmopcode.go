// Package vmopcode — Gopher Workplace challenge.
package vmopcode

// Opcodes for the stack machine. PUSH takes the following word as its operand;
// every other opcode is a single word.
const (
	OpPush = iota
	OpAdd
	OpMul
	OpDup
	OpHalt
)

// VM is a stack machine over a flat program of ints.
type VM struct {
	Prog  []int
	IP    int
	Stack []int
}

// Step executes the instruction at IP and reports whether the VM should keep
// running. It returns false on OpHalt and when IP has run past the program.
// An opcode with too few operands on the stack is a no-op that still advances.
func (v *VM) Step() bool {
	// TODO(candidate): fetch the opcode at IP, advance IP past it (and past
	// PUSH's operand), apply the effect to Stack, and report whether to
	// continue.
	panic("not implemented")
}

// Run steps until the VM halts and returns the value on top of the stack,
// or 0 if the stack is empty.
func (v *VM) Run() int {
	for v.Step() {
	}
	if len(v.Stack) == 0 {
		return 0
	}
	return v.Stack[len(v.Stack)-1]
}
