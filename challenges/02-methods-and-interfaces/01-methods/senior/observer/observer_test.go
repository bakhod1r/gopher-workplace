package observer

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestObserver(t *testing.T) {
	s := &Subject{}

	sum := 0
	s.Attach(func(state int) { sum += state })
	s.Attach(func(state int) { sum += state * 2 })

	s.SetState(10)
	if sum != 30 { // 10 + 20
		t.Errorf("sum = %d, want 30", sum)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardNotifiesObservers(t *testing.T) {
	targets := map[string]bool{"SetState": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "observer.go", nil, 0)
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

	if !seen["observers"] || !seen["state"] {
		t.Logf("WARN: set s.state, then call every observer")
	}
}
