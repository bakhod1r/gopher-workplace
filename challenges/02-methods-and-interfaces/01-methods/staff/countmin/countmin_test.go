package countmin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"testing"
)

func TestCountMin(t *testing.T) {
	s := &Sketch{}
	s.Add("apple")
	s.Add("apple")
	s.Add("bat")

	if got := s.Count("apple"); got != 2 {
		t.Errorf("apple count = %d, want 2", got)
	}
	if got := s.Count("bat"); got != 1 {
		t.Errorf("bat count = %d, want 1", got)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardUsesBothRows(t *testing.T) {
	targets := map[string]bool{"Add": true, "Count": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "countmin.go", nil, 0)
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

	if !seen["row1"] || !seen["row2"] {
		t.Logf("WARN: update and read both rows - a single row has unbounded collision error")
	}
}

// The sketch's guarantee is one-sided: the estimate may exceed the true count
// through collisions, but must never fall below it.
func TestNeverUnderestimates(t *testing.T) {
	s := &Sketch{}
	truth := make(map[string]int)

	for i := 0; i < 5000; i++ {
		item := "k" + strconv.Itoa(i%700)
		s.Add(item)
		truth[item]++
	}

	for item, want := range truth {
		if got := s.Count(item); got < want {
			t.Fatalf("Count(%q) = %d, below the true count %d", item, got, want)
		}
	}
}

// Fixed-size rows: counting must not allocate.
func TestAddIsAllocationFree(t *testing.T) {
	s := &Sketch{}

	allocs := testing.AllocsPerRun(1000, func() {
		s.Add("apple")
		s.Count("apple")
	})
	if allocs != 0 {
		t.Errorf("Add+Count allocated %.0f times per run, want 0", allocs)
	}
}
