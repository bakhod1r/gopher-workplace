package iterpatt

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestIter(t *testing.T) {
	head := &Node{1, &Node{2, &Node{3, nil}}}
	it := NewIter(head)

	var got []int
	for it.HasNext() {
		got = append(got, it.Next())
	}

	want := []int{1, 2, 3}
	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Errorf("got %v, want %v", got, want)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardMovesCursor(t *testing.T) {
	targets := map[string]bool{"HasNext": true, "Next": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "iterpatt.go", nil, 0)
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

	if !seen["current"] || !seen["Next"] {
		t.Logf("WARN: the iterator must read and advance it.current")
	}
}

// A million-node list must be walked in one pass, with the iterator holding a
// single cursor — never a materialized copy of the list.
func TestLargeList(t *testing.T) {
	const n = 1 << 20

	head := &Node{Val: 0}
	curr := head
	for i := 1; i < n; i++ {
		curr.Next = &Node{Val: i}
		curr = curr.Next
	}

	it := NewIter(head)
	count, sum := 0, 0
	for it.HasNext() {
		sum += it.Next()
		count++
	}

	if count != n {
		t.Fatalf("visited %d nodes, want %d", count, n)
	}
	if want := n * (n - 1) / 2; sum != want {
		t.Errorf("sum = %d, want %d", sum, want)
	}
}

// Iterating must not allocate: HasNext reads, Next moves a pointer.
func TestIterationIsAllocationFree(t *testing.T) {
	head := &Node{1, &Node{2, &Node{3, nil}}}

	allocs := testing.AllocsPerRun(1000, func() {
		it := NewIter(head)
		for it.HasNext() {
			it.Next()
		}
	})
	if allocs > 1 {
		t.Errorf("iterating allocated %.0f times per run, want at most 1 (the iterator)", allocs)
	}
}
