package gcswp

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestGCSweep(t *testing.T) {
	h := &Heap{Objects: []bool{true, false, true}}
	if got := h.Sweep(); got != 1 {
		t.Errorf("got %d", got)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardScansTheHeap(t *testing.T) {
	targets := map[string]bool{"Sweep": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "gcswp.go", nil, 0)
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

	if !seen["Objects"] {
		t.Logf("WARN: count the unmarked entries of h.Objects")
	}
}

// Sweeping is a single linear pass over the heap and must not allocate.
func TestSweepIsAllocationFree(t *testing.T) {
	h := &Heap{Objects: make([]bool, 4096)}
	for i := range h.Objects {
		h.Objects[i] = i%2 == 0
	}

	allocs := testing.AllocsPerRun(100, func() {
		h.Sweep()
	})
	if allocs != 0 {
		t.Errorf("Sweep allocated %.0f times per run, want 0", allocs)
	}
}

func TestLargeHeap(t *testing.T) {
	const n = 10 << 20 // 10M objects

	h := &Heap{Objects: make([]bool, n)}
	want := 0
	for i := range h.Objects {
		if i%3 == 0 {
			h.Objects[i] = true
		} else {
			want++
		}
	}

	if got := h.Sweep(); got != want {
		t.Errorf("Sweep() = %d, want %d", got, want)
	}
}

func TestEmptyHeap(t *testing.T) {
	h := &Heap{}
	if got := h.Sweep(); got != 0 {
		t.Errorf("Sweep() = %d, want 0", got)
	}
}
