package actorpatt

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sync"
	"testing"
)

func TestActor(t *testing.T) {
	a := New()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a.Add(1)
		}()
	}
	wg.Wait()

	// Wait for queue to process
	done := make(chan int)
	a.msgs <- func(val *int) { done <- *val }

	if got := <-done; got != 100 {
		t.Errorf("got %d, want 100", got)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardSendsAMessage(t *testing.T) {
	targets := map[string]bool{"Add": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "actorpatt.go", nil, 0)
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

	if !seen["msgs"] {
		t.Logf("WARN: send a closure on a.msgs - touching a.count directly is the race the actor removes")
	}
}
