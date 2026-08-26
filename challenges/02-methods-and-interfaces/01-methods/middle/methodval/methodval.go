// Package methodval — Gopher Workplace challenge.
package methodval

// Greeter holds a greeting prefix.
type Greeter struct {
	Prefix string
}

// Greet returns "<Prefix>, <name>!".
func (g Greeter) Greet(name string) string {
	return g.Prefix + ", " + name + "!"
}

// ApplyMethod takes a Greeter and returns a function that greets by name.
// The returned function is a **bound method value**: it captures the receiver.
//
// Examples:
//
//	fn := ApplyMethod(Greeter{"Hello"})
//	fn("Alice") => "Hello, Alice!"
//	fn("Bob")   => "Hello, Bob!"
func ApplyMethod(g Greeter) func(string) string {
	// TODO(candidate): return the bound method value g.Greet.
	panic("not implemented")
}
