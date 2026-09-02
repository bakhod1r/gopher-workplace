package promisefut

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sync"
	"testing"
	"time"
)

func TestGetBlocksUntilComplete(t *testing.T) {
	f := NewFuture()

	if f.IsDone() {
		t.Error("future should not be done before Complete")
	}

	go func() {
		time.Sleep(10 * time.Millisecond)
		f.Complete(42)
	}()

	if got := f.Get(); got != 42 {
		t.Errorf("Get() = %d, want 42", got)
	}
	if !f.IsDone() {
		t.Error("future should be done after Complete")
	}
}

func TestGetIsRepeatable(t *testing.T) {
	f := NewFuture()
	f.Complete(7)

	for i := 0; i < 3; i++ {
		if got := f.Get(); got != 7 {
			t.Errorf("Get() #%d = %d, want 7", i+1, got)
		}
	}
}

func TestManyWaiters(t *testing.T) {
	f := NewFuture()
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if got := f.Get(); got != 99 {
				t.Errorf("Get() = %d, want 99", got)
			}
		}()
	}

	f.Complete(99)
	wg.Wait()
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardSignalsThroughChannel(t *testing.T) {
	targets := map[string]bool{"Complete": true, "Get": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "promisefut.go", nil, 0)
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

	if !seen["done"] || !seen["close"] {
		t.Logf("WARN: publish the value, then close f.done - that is what releases every waiter")
	}
}
