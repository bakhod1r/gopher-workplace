package exprparser

import "testing"

func TestExpr(t *testing.T) {
	e := &Expr{Val: 5}
	if got := e.Eval(); got != 5 {
		t.Errorf("got %d", got)
	}
}
