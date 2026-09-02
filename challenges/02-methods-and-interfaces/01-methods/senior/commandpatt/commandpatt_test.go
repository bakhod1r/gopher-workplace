package commandpatt

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestCommand(t *testing.T) {
	inv := &Invoker{}
	var x int

	inv.Add(func() { x += 5 })
	inv.Add(func() { x *= 2 })

	inv.ExecuteAll()
	if x != 10 {
		t.Errorf("x = %d, want 10", x)
	}

	if len(inv.commands) != 0 {
		t.Error("queue not cleared")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardRunsAndClearsQueue(t *testing.T) {
	targets := map[string]bool{"ExecuteAll": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "commandpatt.go", nil, 0)
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

	if !seen["commands"] {
		t.Logf("WARN: run everything in inv.commands, then clear the queue")
	}
}
