package backoff

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
	"time"
)

func TestBackoff(t *testing.T) {
	b := New(5 * time.Second)

	wants := []time.Duration{
		1 * time.Second,
		2 * time.Second,
		4 * time.Second,
		5 * time.Second,
		5 * time.Second,
	}

	for i, want := range wants {
		if got := b.Next(); got != want {
			t.Errorf("call %d: got %v, want %v", i, got, want)
		}
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardUsesStoredDelay(t *testing.T) {
	targets := map[string]bool{"Next": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "backoff.go", nil, 0)
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

	if !seen["current"] || !seen["max"] {
		t.Logf("WARN: derive the delay from b.current and clamp with b.max - don't hand-type the sequence")
	}
}
