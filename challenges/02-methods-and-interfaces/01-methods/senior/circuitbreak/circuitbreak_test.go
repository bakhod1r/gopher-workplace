package circuitbreak

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestBreaker(t *testing.T) {
	b := &Breaker{Threshold: 2}
	failFn := func() error { return errors.New("fail") }
	okFn := func() error { return nil }

	if err := b.Call(failFn); err == nil {
		t.Error("expected error")
	}
	if b.IsOpen {
		t.Error("should not be open yet")
	}

	if err := b.Call(failFn); err == nil {
		t.Error("expected error")
	}
	if !b.IsOpen {
		t.Error("should be open now")
	}

	err := b.Call(okFn)
	if err == nil || err.Error() != "circuit open" {
		t.Errorf("expected circuit open error, got %v", err)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardChecksOpenState(t *testing.T) {
	targets := map[string]bool{"Call": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "circuitbreak.go", nil, 0)
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

	if !seen["IsOpen"] || !seen["ConsecutiveFails"] {
		t.Logf("WARN: the breaker must consult IsOpen before calling fn and count ConsecutiveFails")
	}
}
