package visitorpatt

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestVisitor(t *testing.T) {
	root := &Node{
		Val:   1,
		Left:  &Node{Val: 2},
		Right: &Node{Val: 3},
	}

	sum := 0
	root.Accept(func(v int) { sum += v })

	if sum != 6 {
		t.Errorf("sum = %d, want 6", sum)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardVisitsAndRecurses(t *testing.T) {
	targets := map[string]bool{"Accept": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "visitorpatt.go", nil, 0)
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

	if !seen["visitor"] || !seen["Left"] || !seen["Right"] {
		t.Logf("WARN: call the visitor and recurse into both Left and Right")
	}
}
