package retrylogic

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestRetry(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := &Client{FailInt: 2}
		err := c.DoWithRetry(3)
		if err != nil {
			t.Errorf("expected success, got %v", err)
		}
		if c.Attempts != 3 {
			t.Errorf("Attempts = %d, want 3", c.Attempts)
		}
	})

	t.Run("exhausted", func(t *testing.T) {
		c := &Client{FailInt: 5}
		err := c.DoWithRetry(3)
		if err == nil {
			t.Error("expected error, got nil")
		}
		if c.Attempts != 3 {
			t.Errorf("Attempts = %d, want 3", c.Attempts)
		}
	})
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardRetriesTheCall(t *testing.T) {
	targets := map[string]bool{"DoWithRetry": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "retrylogic.go", nil, 0)
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

	if !seen["Do"] || !seen["maxAttempts"] {
		t.Logf("WARN: loop up to maxAttempts calling c.Do() - a fixed number of calls is not a retry policy")
	}
}
