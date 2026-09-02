package stmtransact

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sync"
	"testing"
)

func TestTxAppliesFunction(t *testing.T) {
	tv := NewTVar(5)

	if got := tv.Tx(func(v int) int { return v * 2 }); got != 10 {
		t.Errorf("Tx returned %d, want 10", got)
	}
	if got, _ := tv.Read(); got != 10 {
		t.Errorf("value = %d, want 10", got)
	}
}

func TestTxBumpsVersion(t *testing.T) {
	tv := NewTVar(0)
	_, v0 := tv.Read()

	tv.Tx(func(v int) int { return v + 1 })

	_, v1 := tv.Read()
	if v1 != v0+1 {
		t.Errorf("version = %d, want %d", v1, v0+1)
	}
}

func TestTxRetriesUnderContention(t *testing.T) {
	tv := NewTVar(0)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tv.Tx(func(v int) int { return v + 1 })
		}()
	}
	wg.Wait()

	if got, _ := tv.Read(); got != 100 {
		t.Errorf("value = %d, want 100 — a lost update means a commit was not retried", got)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardRetriesOnConflict(t *testing.T) {
	targets := map[string]bool{"Tx": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "stmtransact.go", nil, 0)
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

	if !seen["Read"] || !seen["Commit"] {
		t.Logf("WARN: Read, compute, Commit - and re-Read on a rejected commit")
	}
}
