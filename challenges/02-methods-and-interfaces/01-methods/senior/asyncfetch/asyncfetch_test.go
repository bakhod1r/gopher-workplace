package asyncfetch

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
	"time"
)

func TestFetchAsync(t *testing.T) {
	f := &Fetcher{}
	done := f.FetchAsync("abc")

	select {
	case <-done:
		if f.Result != "data: abc" {
			t.Errorf("Result = %q, want %q", f.Result, "data: abc")
		}
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for FetchAsync")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardFetchesInGoroutine(t *testing.T) {
	targets := map[string]bool{"FetchAsync": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "asyncfetch.go", nil, 0)
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

	if !seen["Fetch"] || !seen["close"] {
		t.Logf("WARN: launch Fetch in a goroutine and close the returned channel when it is done")
	}
}
