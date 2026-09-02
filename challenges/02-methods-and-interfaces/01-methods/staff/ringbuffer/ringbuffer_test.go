package ringbuffer

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestPushUntilFull(t *testing.T) {
	r := New(2)

	if err := r.Push(1); err != nil {
		t.Fatalf("Push(1) = %v, want nil", err)
	}
	if err := r.Push(2); err != nil {
		t.Fatalf("Push(2) = %v, want nil", err)
	}
	if err := r.Push(3); !errors.Is(err, ErrFull) {
		t.Errorf("Push on full = %v, want ErrFull", err)
	}
	if r.Len() != 2 {
		t.Errorf("Len() = %d, want 2 — a rejected Push must change nothing", r.Len())
	}
}

func TestPopFIFOAndEmpty(t *testing.T) {
	r := New(3)
	r.Push(1)
	r.Push(2)

	for _, want := range []int{1, 2} {
		got, err := r.Pop()
		if err != nil || got != want {
			t.Fatalf("Pop() = (%d, %v), want (%d, nil)", got, err, want)
		}
	}

	if _, err := r.Pop(); !errors.Is(err, ErrEmpty) {
		t.Errorf("Pop on empty = %v, want ErrEmpty", err)
	}
}

func TestWrapAround(t *testing.T) {
	r := New(3)
	r.Push(1)
	r.Push(2)
	r.Push(3)
	r.Pop() // 1
	r.Pop() // 2

	if err := r.Push(4); err != nil { // reuses slot 0
		t.Fatalf("Push after wrap = %v", err)
	}
	if err := r.Push(5); err != nil { // reuses slot 1
		t.Fatalf("Push after wrap = %v", err)
	}

	for _, want := range []int{3, 4, 5} {
		got, err := r.Pop()
		if err != nil || got != want {
			t.Errorf("Pop() = (%d, %v), want (%d, nil)", got, err, want)
		}
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardMovesCursors(t *testing.T) {
	targets := map[string]bool{"Push": true, "Pop": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "ringbuffer.go", nil, 0)
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

	if !seen["head"] || !seen["tail"] || !seen["size"] {
		t.Logf("WARN: move head/tail and keep size in step - that is the ring")
	}
}

// New allocates the backing slice once; steady-state use must allocate nothing.
func TestPushPopAreAllocationFree(t *testing.T) {
	r := New(8)

	allocs := testing.AllocsPerRun(1000, func() {
		r.Push(1)
		r.Pop()
	})
	if allocs != 0 {
		t.Errorf("Push+Pop allocated %.0f times per run, want 0", allocs)
	}
}

// A rejected Push must be free of side effects even after many wraps.
func TestFullBufferStaysConsistent(t *testing.T) {
	r := New(4)
	for i := 0; i < 4; i++ {
		r.Push(i)
	}

	for i := 0; i < 1000; i++ {
		if err := r.Push(99); err == nil {
			t.Fatal("Push on a full buffer returned nil error")
		}
	}
	if r.Len() != 4 {
		t.Fatalf("Len() = %d, want 4", r.Len())
	}
	for want := 0; want < 4; want++ {
		if got, err := r.Pop(); err != nil || got != want {
			t.Fatalf("Pop() = (%d, %v), want (%d, nil)", got, err, want)
		}
	}
}
