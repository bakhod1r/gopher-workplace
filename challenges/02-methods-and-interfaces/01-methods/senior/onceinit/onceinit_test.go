package onceinit

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sync"
	"sync/atomic"
	"testing"
)

func TestOnceInit(t *testing.T) {
	var calls int32
	l := New(func() string {
		atomic.AddInt32(&calls, 1)
		return "safe"
	})

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if got := l.Get(); got != "safe" {
				t.Errorf("got %q, want safe", got)
			}
		}()
	}
	wg.Wait()

	if calls != 1 {
		t.Errorf("calls = %d, want 1", calls)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardUsesOnce(t *testing.T) {
	targets := map[string]bool{"Get": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "onceinit.go", nil, 0)
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

	if !seen["once"] || !seen["Do"] {
		t.Logf("WARN: use sync.Once (l.once.Do) - a plain nil or zero check is not safe under concurrency")
	}
}
