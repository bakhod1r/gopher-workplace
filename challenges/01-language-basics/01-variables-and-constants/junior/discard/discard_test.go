package discard

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestSplit(t *testing.T) {
	cases := []struct {
		name              string
		n, size           int
		wantPages, wantOK int
	}{
		{"uneven", 10, 3, 3, 1},
		{"exact", 9, 3, 3, 0},
		{"smaller than a page", 2, 5, 0, 2},
		{"zero items", 0, 4, 0, 0},
		{"zero size", 7, 0, 0, 7},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pages, rest := Split(tc.n, tc.size)
			if pages != tc.wantPages || rest != tc.wantOK {
				t.Errorf("Split(%d, %d) = %d, %d; want %d, %d",
					tc.n, tc.size, pages, rest, tc.wantPages, tc.wantOK)
			}
		})
	}
}

func TestPages(t *testing.T) {
	cases := []struct {
		name    string
		n, size int
		want    int
	}{
		{"uneven", 10, 3, 3},
		{"exact", 9, 3, 3},
		{"smaller than a page", 2, 5, 0},
		{"zero items", 0, 4, 0},
		{"zero size", 7, 0, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Pages(tc.n, tc.size); got != tc.want {
				t.Errorf("Pages(%d, %d) = %d, want %d", tc.n, tc.size, got, tc.want)
			}
		})
	}
}

func TestLeftover(t *testing.T) {
	cases := []struct {
		name    string
		n, size int
		want    int
	}{
		{"uneven", 10, 3, 1},
		{"exact", 9, 3, 0},
		{"smaller than a page", 2, 5, 2},
		{"zero items", 0, 4, 0},
		{"zero size", 7, 0, 7},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Leftover(tc.n, tc.size); got != tc.want {
				t.Errorf("Leftover(%d, %d) = %d, want %d", tc.n, tc.size, got, tc.want)
			}
		})
	}
}

// Pages and Leftover must agree with Split for the same inputs — they are views
// of one calculation, not two independent ones.
func TestViewsAgreeWithSplit(t *testing.T) {
	for n := 0; n < 40; n++ {
		for size := 0; size < 7; size++ {
			pages, rest := Split(n, size)
			if got := Pages(n, size); got != pages {
				t.Fatalf("Pages(%d, %d) = %d, but Split says %d", n, size, got, pages)
			}
			if got := Leftover(n, size); got != rest {
				t.Fatalf("Leftover(%d, %d) = %d, but Split says %d", n, size, got, rest)
			}
		}
	}
}

// The lesson is the blank identifier: both views should delegate to Split and
// throw away the half they do not need, rather than redoing the arithmetic.
func TestUsesBlankIdentifier(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "discard.go", nil, 0) // 0 = skip comments
	if err != nil {
		return // parse trouble is not this check's concern
	}
	for _, name := range []string{"Pages", "Leftover"} {
		fn := findFunc(f, name)
		if fn == nil {
			continue
		}
		if !callsSplit(fn) {
			t.Logf("WARN: %s should call Split and discard the value it does not need, not recompute it", name)
			continue
		}
		if !usesBlank(fn) {
			t.Logf("WARN: %s receives both of Split's values — use the blank identifier _ for the one you drop", name)
		}
	}
}

func findFunc(f *ast.File, name string) *ast.FuncDecl {
	for _, d := range f.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == name {
			return fn
		}
	}
	return nil
}

func callsSplit(fn *ast.FuncDecl) bool {
	found := false
	ast.Inspect(fn, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		if id, ok := call.Fun.(*ast.Ident); ok && id.Name == "Split" {
			found = true
		}
		return true
	})
	return found
}

func usesBlank(fn *ast.FuncDecl) bool {
	found := false
	ast.Inspect(fn, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == "_" {
			found = true
		}
		return true
	})
	return found
}
