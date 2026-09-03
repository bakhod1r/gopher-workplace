package uintptrhold

import (
	"go/ast"
	"go/parser"
	"go/token"
	"runtime"
	"testing"
)

func TestSecondWord(t *testing.T) {
	if got := SecondWord(&Pair{A: 1, B: 2}); got != 2 {
		t.Errorf("SecondWord = %d, want 2", got)
	}
	if got := SecondWord(&Pair{}); got != 0 {
		t.Errorf("SecondWord = %d, want 0", got)
	}
	if got := SecondWord(&Pair{A: 9, B: -7}); got != -7 {
		t.Errorf("SecondWord = %d, want -7", got)
	}
}

func TestSecondWordUnderAllocationPressure(t *testing.T) {
	for i := 0; i < 2000; i++ {
		p := &Pair{A: int64(i), B: int64(i * 2)}
		if got := SecondWord(p); got != int64(i*2) {
			t.Fatalf("iteration %d: SecondWord = %d, want %d", i, got, i*2)
		}
		if i%100 == 0 {
			runtime.GC()
		}
		runtime.KeepAlive(p)
	}
}

func TestSecondWordUsesUnsafeAdd(t *testing.T) {
	// A moving collector would invalidate a stored uintptr, and no test can
	// provoke that on demand -- so the rule is checked in the source instead.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "uintptrhold.go", nil, 0)
	if err != nil {
		t.Skipf("cannot parse the source: %v", err)
	}
	usesAdd := false
	holdsUintptr := false
	ast.Inspect(f, func(n ast.Node) bool {
		if sel, ok := n.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok && id.Name == "unsafe" && sel.Sel.Name == "Add" {
				usesAdd = true
			}
		}
		// A uintptr conversion assigned to a variable outlives its expression.
		record := func(rhs []ast.Expr) {
			for _, e := range rhs {
				ast.Inspect(e, func(m ast.Node) bool {
					c, ok := m.(*ast.CallExpr)
					if !ok {
						return true
					}
					if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "uintptr" {
						holdsUintptr = true
					}
					return true
				})
			}
		}
		switch s := n.(type) {
		case *ast.AssignStmt:
			record(s.Rhs)
		case *ast.ValueSpec:
			record(s.Values)
		}
		return true
	})
	if holdsUintptr {
		t.Error("a uintptr conversion is stored in a variable: the address goes stale outside its expression")
	}
	if !usesAdd {
		t.Error("the offset must be applied with unsafe.Add, which keeps the result a pointer")
	}
}
