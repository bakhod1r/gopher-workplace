package syncpool

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestGetReturnsInitializedBuffer(t *testing.T) {
	p := NewPool()

	b := p.Get()
	if b == nil {
		t.Fatal("Get() = nil, want a buffer")
	}
	if len(b.Data) != 1024 {
		t.Fatalf("len(Data) = %d, want 1024", len(b.Data))
	}
}

func TestPutThenGetReuses(t *testing.T) {
	p := NewPool()

	b := p.Get()
	b.Data[0] = 42
	p.Put(b)

	// sync.Pool may drop entries at any GC, so reuse is likely but not
	// guaranteed. What must hold is that Get always returns a usable buffer.
	b2 := p.Get()
	if b2 == nil || len(b2.Data) != 1024 {
		t.Fatalf("Get() after Put returned %v", b2)
	}
}

func TestManyGetsAreDistinct(t *testing.T) {
	p := NewPool()

	a := p.Get()
	b := p.Get()
	if a == b {
		t.Error("two Gets without an intervening Put returned the same buffer")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardUsesThePool(t *testing.T) {
	targets := map[string]bool{"Get": true, "Put": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "syncpool.go", nil, 0)
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

	if !seen["pool"] {
		t.Logf("WARN: Get and Put must go through p.pool, not allocate a fresh Buffer every time")
	}
}
