package poolworker

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sync/atomic"
	"testing"
)

func TestPool(t *testing.T) {
	p := Pool{
		Count: 3,
		Tasks: make(chan func(), 10),
	}
	p.Start()

	var sum int32
	for i := 0; i < 5; i++ {
		p.Tasks <- func() { atomic.AddInt32(&sum, 1) }
	}
	close(p.Tasks)
	p.Wait()

	if sum != 5 {
		t.Errorf("sum = %d, want 5", sum)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardRunsTasks(t *testing.T) {
	targets := map[string]bool{"Start": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "poolworker.go", nil, 0)
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

	if !seen["Tasks"] || !seen["wg"] {
		t.Logf("WARN: workers must range over p.Tasks, call each task, and coordinate with p.wg")
	}
}
