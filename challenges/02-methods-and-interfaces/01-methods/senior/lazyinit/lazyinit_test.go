package lazyinit

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestLazy(t *testing.T) {
	calls := 0
	l := New(func() string {
		calls++
		return "heavy"
	})

	if calls != 0 {
		t.Error("init called early")
	}

	if got := l.String(); got != "heavy" || calls != 1 {
		t.Errorf("first got=%q, calls=%d", got, calls)
	}

	if got := l.String(); got != "heavy" || calls != 1 {
		t.Errorf("second got=%q, calls=%d", got, calls)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardCallsInit(t *testing.T) {
	targets := map[string]bool{"String": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "lazyinit.go", nil, 0)
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

	if !seen["init"] {
		t.Logf("WARN: produce the value by calling l.init(), don't return a hardcoded string")
	}
}
