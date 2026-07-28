package temperature

import (
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"testing"
)

// TestNoHardcodedOutput checks the task's "compute it, don't hardcode" rule: a
// real conversion must use the input c, whereas a lookup table or hand-typed
// return values ignore it. It does NOT fail the submission — it emits a `WARN:`
// line (surfaced by the runner) when CToF never references its parameter c.
func TestNoHardcodedOutput(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "temperature.go", nil, 0)
	if err != nil {
		return
	}
	usesC := false
	ast.Inspect(f, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok || fn.Name.Name != "CToF" || fn.Body == nil {
			return true
		}
		ast.Inspect(fn.Body, func(m ast.Node) bool {
			if id, ok := m.(*ast.Ident); ok && id.Name == "c" {
				usesC = true
			}
			return true
		})
		return false
	})
	if !usesC {
		t.Logf("WARN: CToF ignores its input c — compute F from c, don't hardcode the results")
	}
}

func TestCToF(t *testing.T) {
	cases := []struct {
		name string
		c    float64
		want float64
	}{
		{"freezing", 0, 32},
		{"boiling", 100, 212},
		{"crossover", -40, -40},
		{"body", 37, 98.6},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CToF(tc.c); math.Abs(got-tc.want) > 1e-9 {
				t.Errorf("CToF(%g) = %g, want %g", tc.c, got, tc.want)
			}
		})
	}
}
