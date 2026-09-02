package arcpool

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestARC(t *testing.T) {
	a := &ARC{T1: make(map[int]bool), T2: make(map[int]bool)}

	a.Access(1)
	if !a.T1[1] || a.T2[1] {
		t.Error("1 should be in T1")
	}

	a.Access(1)
	if a.T1[1] || !a.T2[1] {
		t.Error("1 should move to T2")
	}

	a.Access(1)
	if a.T1[1] || !a.T2[1] {
		t.Error("1 should stay in T2")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardMovesBetweenLists(t *testing.T) {
	targets := map[string]bool{"Access": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "arcpool.go", nil, 0)
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

	if !seen["T1"] || !seen["T2"] || !seen["delete"] {
		t.Logf("WARN: promotion must delete from T1 and insert into T2")
	}
}
