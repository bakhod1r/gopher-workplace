// Package jitsim — Gopher Workplace challenge.
package jitsim

// Tier reports how an operation was executed.
type Tier string

const (
	// Interpreted means the op ran through the interpreter.
	Interpreted Tier = "interp"
	// Compiled means the op was hot enough to run as compiled code.
	Compiled Tier = "jit"
)

// JIT counts how often each operation runs and promotes an operation to
// compiled code once it has been seen Threshold times.
type JIT struct {
	Threshold int
	counts    map[string]int
	compiled  map[string]bool
}

// New returns a JIT that compiles an op on its threshold-th execution.
func New(threshold int) *JIT {
	return &JIT{
		Threshold: threshold,
		counts:    make(map[string]int),
		compiled:  make(map[string]bool),
	}
}

// Execute runs op and reports which tier served it. The execution that reaches
// Threshold is itself served by the compiled tier, as are all later ones.
func (j *JIT) Execute(op string) Tier {
	// TODO(candidate): count the execution; once the count reaches Threshold,
	// mark the op compiled and report Compiled — otherwise Interpreted.
	panic("not implemented")
}

// IsCompiled reports whether op has been promoted.
func (j *JIT) IsCompiled(op string) bool {
	return j.compiled[op]
}
