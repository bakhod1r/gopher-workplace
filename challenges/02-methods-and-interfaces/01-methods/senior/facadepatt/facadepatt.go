// Package facadepatt — Gopher Workplace challenge.
package facadepatt

type Sub1 struct{}

func (Sub1) Op1() string { return "1" }

type Sub2 struct{}

func (Sub2) Op2() string { return "2" }

// Facade simplifies Sub1 and Sub2.
type Facade struct {
	s1 Sub1
	s2 Sub2
}

// Operation combines both subsystems.
func (f *Facade) Operation() string {
	// TODO(candidate): return f.s1.Op1() + "+" + f.s2.Op2()
	panic("not implemented")
}
