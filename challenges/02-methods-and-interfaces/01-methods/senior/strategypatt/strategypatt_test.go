package strategypatt

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestStrategy(t *testing.T) {
	c := &Context{Data: []int{1, 2, 3}}

	double := func(x int) int { return x * 2 }
	c.Process(double)

	if !reflect.DeepEqual(c.Data, []int{2, 4, 6}) {
		t.Errorf("Process double = %v", c.Data)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardAppliesStrategy(t *testing.T) {
	targets := map[string]bool{"Process": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "strategypatt.go", nil, 0)
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

	if !seen["strategy"] || !seen["Data"] {
		t.Logf("WARN: apply the strategy parameter to c.Data - don't compute the expected values directly")
	}
}

func double(x int) int { return x * 2 }

// Process rewrites the slice it was given: applying a strategy to a million
// elements must not allocate a second copy.
func TestProcessIsAllocationFree(t *testing.T) {
	c := &Context{Data: make([]int, 1<<20)}

	allocs := testing.AllocsPerRun(5, func() {
		c.Process(double)
	})
	if allocs != 0 {
		t.Errorf("Process allocated %.0f times per run, want 0 — it must work in place", allocs)
	}
}

func TestLargeInput(t *testing.T) {
	const n = 1 << 20

	data := make([]int, n)
	for i := range data {
		data[i] = i
	}
	c := &Context{Data: data}

	c.Process(double)

	if len(c.Data) != n {
		t.Fatalf("len = %d, want %d", len(c.Data), n)
	}
	if c.Data[0] != 0 || c.Data[1] != 2 || c.Data[n-1] != (n-1)*2 {
		t.Errorf("unexpected values: %d %d %d", c.Data[0], c.Data[1], c.Data[n-1])
	}
}
