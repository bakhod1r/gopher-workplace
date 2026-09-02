package builderpatt

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestBuilder(t *testing.T) {
	req := NewBuilder().
		Method("GET").
		URL("/api").
		Auth("token123").
		Build()

	if req.Method != "GET" || req.URL != "/api" || req.Auth != "token123" {
		t.Errorf("Build failed: %+v", req)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardReturnsReceiver(t *testing.T) {
	targets := map[string]bool{"URL": true, "Auth": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "builderpatt.go", nil, 0)
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

	if !seen["req"] {
		t.Logf("WARN: each setter writes into b.req and returns b so the chain continues")
	}
}
