package io_uring

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestSubmitCompleteFIFO(t *testing.T) {
	r := &Ring{}

	for _, op := range []int{10, 20, 30} {
		if !r.Submit(op) {
			t.Fatalf("Submit(%d) = false, want true", op)
		}
	}
	if r.Len() != 3 {
		t.Errorf("Len() = %d, want 3", r.Len())
	}

	for _, want := range []int{10, 20, 30} {
		got, ok := r.Complete()
		if !ok || got != want {
			t.Errorf("Complete() = (%d, %v), want (%d, true)", got, ok, want)
		}
	}
	if r.Len() != 0 {
		t.Errorf("Len() = %d, want 0", r.Len())
	}
}

func TestFullAndEmpty(t *testing.T) {
	r := &Ring{}

	for i := 0; i < ringSize; i++ {
		if !r.Submit(i) {
			t.Fatalf("Submit(%d) = false while ring had room", i)
		}
	}
	if r.Submit(99) {
		t.Error("Submit on a full ring = true, want false")
	}

	for i := 0; i < ringSize; i++ {
		if _, ok := r.Complete(); !ok {
			t.Fatalf("Complete() = false while ring had entries")
		}
	}
	if _, ok := r.Complete(); ok {
		t.Error("Complete on an empty ring = true, want false")
	}
}

func TestWrapAround(t *testing.T) {
	r := &Ring{}

	// Fill, drain most of it, then refill so the indices wrap.
	for i := 0; i < ringSize; i++ {
		r.Submit(i)
	}
	for i := 0; i < ringSize-1; i++ {
		r.Complete()
	}
	for i := 100; i < 100+ringSize-1; i++ {
		if !r.Submit(i) {
			t.Fatalf("Submit(%d) = false after wrap", i)
		}
	}

	want := []int{ringSize - 1, 100, 101}
	for _, w := range want {
		got, ok := r.Complete()
		if !ok || got != w {
			t.Errorf("Complete() = (%d, %v), want (%d, true)", got, ok, w)
		}
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardMovesCursors(t *testing.T) {
	targets := map[string]bool{"Submit": true, "Complete": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "io_uring.go", nil, 0)
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

	if !seen["head"] || !seen["tail"] || !seen["count"] {
		t.Logf("WARN: move head/tail and keep count in step - that is the ring")
	}
}

// The ring is fixed-capacity by construction: after New, a Submit/Complete
// round trip must not allocate.
func TestSubmitCompleteAreAllocationFree(t *testing.T) {
	r := &Ring{}

	allocs := testing.AllocsPerRun(1000, func() {
		r.Submit(1)
		r.Complete()
	})
	if allocs != 0 {
		t.Errorf("Submit+Complete allocated %.0f times per run, want 0", allocs)
	}
}

// Cursors must survive far more traversals of the ring than it has slots.
func TestLongRunKeepsOrder(t *testing.T) {
	r := &Ring{}

	for i := 0; i < 100000; i++ {
		if !r.Submit(i) {
			t.Fatalf("Submit(%d) rejected on an empty ring", i)
		}
		got, ok := r.Complete()
		if !ok || got != i {
			t.Fatalf("Complete() = (%d, %v), want (%d, true)", got, ok, i)
		}
	}
	if r.Len() != 0 {
		t.Errorf("Len() = %d, want 0", r.Len())
	}
}
