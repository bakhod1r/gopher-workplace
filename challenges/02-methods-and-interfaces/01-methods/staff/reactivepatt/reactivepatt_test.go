package reactivepatt

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestStream(t *testing.T) {
	s := &Stream{Data: []int{1, 2, 3, 4}}

	s.Filter(func(x int) bool { return x%2 == 0 }).
		Map(func(x int) int { return x * 10 })

	if !reflect.DeepEqual(s.Data, []int{20, 40}) {
		t.Errorf("got %v", s.Data)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardFiltersAndMaps(t *testing.T) {
	targets := map[string]bool{"Filter": true, "Map": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "reactivepatt.go", nil, 0)
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

	if !seen["fn"] || !seen["Data"] {
		t.Logf("WARN: apply fn to s.Data in place and return the receiver")
	}
}

func isEven(x int) bool  { return x%2 == 0 }
func timesTen(x int) int { return x * 10 }

// Both operators rewrite the existing backing array, so a pipeline over an
// already-allocated stream must not allocate at all.
func TestPipelineIsAllocationFree(t *testing.T) {
	data := make([]int, 1024)
	s := &Stream{}

	allocs := testing.AllocsPerRun(100, func() {
		for i := range data {
			data[i] = i
		}
		s.Data = data
		s.Filter(isEven).Map(timesTen)
	})
	if allocs != 0 {
		t.Errorf("Filter+Map allocated %.0f times per run, want 0", allocs)
	}
}

func TestLargeStream(t *testing.T) {
	const n = 1 << 20
	data := make([]int, n)
	for i := range data {
		data[i] = i
	}
	s := &Stream{Data: data}

	s.Filter(isEven).Map(timesTen)

	if len(s.Data) != n/2 {
		t.Fatalf("len = %d, want %d", len(s.Data), n/2)
	}
	if s.Data[0] != 0 || s.Data[1] != 20 || s.Data[n/2-1] != (n-2)*10 {
		t.Errorf("unexpected values: %d %d %d", s.Data[0], s.Data[1], s.Data[n/2-1])
	}
}
