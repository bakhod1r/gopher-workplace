package templatemeth

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestTemplate(t *testing.T) {
	tmpl := &Template{impl: MyTask{}}
	if got := tmpl.Run(); got != "a-b" {
		t.Errorf("Run() = %q, want a-b", got)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardCallsBothSteps(t *testing.T) {
	targets := map[string]bool{"Run": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "templatemeth.go", nil, 0)
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

	if !seen["DoStep1"] || !seen["DoStep2"] {
		t.Logf("WARN: call the injected steps (DoStep1, DoStep2) - the template method must not hardcode their output")
	}
}
