package statepattern

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestState(t *testing.T) {
	d := &Document{State: Draft}
	d.Publish()
	if d.State != Moderation {
		t.Errorf("expected Moderation")
	}
	d.Publish()
	if d.State != Published {
		t.Errorf("expected Published")
	}
	d.Publish()
	if d.State != Published {
		t.Errorf("expected Published")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardBranchesOnState(t *testing.T) {
	targets := map[string]bool{"Publish": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "statepattern.go", nil, 0)
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

	if !seen["State"] {
		t.Logf("WARN: branch on the current State - arithmetic on the enum walks past the terminal state")
	}
}
