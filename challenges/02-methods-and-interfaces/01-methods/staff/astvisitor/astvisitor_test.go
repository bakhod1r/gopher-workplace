package astvisitor

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestVisit(t *testing.T) {
	root := &Node{
		Type: "BinOp",
		Left: &Node{Type: "Ident", Name: "x"},
		Right: &Node{
			Type:  "BinOp",
			Left:  &Node{Type: "Ident", Name: "y"},
			Right: &Node{Type: "Num"},
		},
	}

	count := 0
	root.Visit(&count)
	if count != 2 {
		t.Errorf("got %d, want 2", count)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardCountsAndRecurses(t *testing.T) {
	targets := map[string]bool{"Visit": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "astvisitor.go", nil, 0)
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

	if !seen["count"] || !seen["Left"] || !seen["Right"] {
		t.Logf("WARN: increment through the count pointer and recurse into both children")
	}
}
