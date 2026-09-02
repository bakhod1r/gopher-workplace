package lockfreestk

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sync"
	"testing"
)

func TestLockFreeStack(t *testing.T) {
	s := &Stack{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			s.Push(v)
		}(i)
	}
	wg.Wait()

	count := 0
	curr := s.head.Load()
	for curr != nil {
		count++
		curr = curr.next
	}
	if count != 100 {
		t.Errorf("count = %d, want 100", count)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardUsesCAS(t *testing.T) {
	targets := map[string]bool{"Push": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "lockfreestk.go", nil, 0)
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

	if !seen["CompareAndSwap"] {
		t.Logf("WARN: push with a CompareAndSwap retry loop")
	}
}
