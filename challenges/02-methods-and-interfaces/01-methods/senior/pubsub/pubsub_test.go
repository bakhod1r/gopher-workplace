package pubsub

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestPubSub(t *testing.T) {
	ps := New()
	ch1 := ps.Subscribe("news")
	ch2 := ps.Subscribe("news")

	ps.Publish("news", "hello")

	if got := <-ch1; got != "hello" {
		t.Errorf("ch1 got %q", got)
	}
	if got := <-ch2; got != "hello" {
		t.Errorf("ch2 got %q", got)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardUsesReadLock(t *testing.T) {
	targets := map[string]bool{"Publish": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "pubsub.go", nil, 0)
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

	if !seen["RLock"] {
		t.Logf("WARN: read the subscriber map under RLock - Lock serializes publishers needlessly, no lock at all is a race")
	}
}
