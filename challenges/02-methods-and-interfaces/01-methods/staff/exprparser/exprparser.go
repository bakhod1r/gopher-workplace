// Package exprparser — Gopher Workplace challenge.
package exprparser

// Kind tags an expression node.
type Kind int

const (
	Num Kind = iota
	Add
	Mul
	Neg
)

// Expr is a node in an arithmetic expression tree.
//
//	Num — a literal; Val holds the number.
//	Add — Left + Right.
//	Mul — Left * Right.
//	Neg — negation of Left; Right is unused.
type Expr struct {
	Kind        Kind
	Val         int
	Left, Right *Expr
}

// Eval evaluates the expression tree rooted at e.
// A nil expression evaluates to 0.
func (e *Expr) Eval() int {
	// TODO(candidate): evaluate by Kind, recursing into Left and Right.
	panic("not implemented")
}
