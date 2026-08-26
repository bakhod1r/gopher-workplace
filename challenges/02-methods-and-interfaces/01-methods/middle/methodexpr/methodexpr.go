// Package methodexpr — Gopher Workplace challenge.
package methodexpr

// Adder can add a value to its base.
type Adder struct {
	Base int
}

// Add returns Base + n.
func (a Adder) Add(n int) int {
	return a.Base + n
}

// CallExpr takes a method expression and applies it to the given Adder and
// argument. A method expression has the receiver as the first parameter:
//
//	Adder.Add  has type  func(Adder, int) int
//
// Examples:
//
//	CallExpr(Adder.Add, Adder{10}, 5) => 15
func CallExpr(fn func(Adder, int) int, a Adder, n int) int {
	// TODO(candidate): call fn with a and n.
	panic("not implemented")
}
