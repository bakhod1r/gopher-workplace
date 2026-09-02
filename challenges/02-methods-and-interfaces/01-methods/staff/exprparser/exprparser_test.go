package exprparser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestEvalLiteral(t *testing.T) {
	e := &Expr{Kind: Num, Val: 5}
	if got := e.Eval(); got != 5 {
		t.Errorf("Eval() = %d, want 5", got)
	}
}

func TestEvalNested(t *testing.T) {
	// (2 + 3) * 4
	e := &Expr{
		Kind: Mul,
		Left: &Expr{
			Kind:  Add,
			Left:  &Expr{Kind: Num, Val: 2},
			Right: &Expr{Kind: Num, Val: 3},
		},
		Right: &Expr{Kind: Num, Val: 4},
	}
	if got := e.Eval(); got != 20 {
		t.Errorf("Eval() = %d, want 20", got)
	}
}

func TestEvalNegAndNil(t *testing.T) {
	// -(7 * 2)
	e := &Expr{
		Kind: Neg,
		Left: &Expr{
			Kind:  Mul,
			Left:  &Expr{Kind: Num, Val: 7},
			Right: &Expr{Kind: Num, Val: 2},
		},
	}
	if got := e.Eval(); got != -14 {
		t.Errorf("Eval() = %d, want -14", got)
	}

	var nilExpr *Expr
	if got := nilExpr.Eval(); got != 0 {
		t.Errorf("nil Eval() = %d, want 0", got)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardEvaluatesRecursively(t *testing.T) {
	targets := map[string]bool{"Eval": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "exprparser.go", nil, 0)
	if err != nil {
		return // parse trouble is not this check's concern
	}

	seen := map[string]bool{}
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil || !targets[fn.Name.Name] {
			continue
		}
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.SelectorExpr:
				seen[v.Sel.Name] = true
			case *ast.Ident:
				seen[v.Name] = true
			}
			return true
		})
	}

	if !seen["Eval"] || !seen["Kind"] {
		t.Logf("WARN: evaluate by node Kind, recursing with Eval - don't special-case the test's tree")
	}
}
