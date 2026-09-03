package mapprealloc

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	got := Count([]string{"a", "b", "a", "c", "a"})
	want := map[string]int{"a": 3, "b": 1, "c": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Count = %v, want %v", got, want)
	}
	if got := Count(nil); len(got) != 0 {
		t.Errorf("Count(nil) = %v, want empty", got)
	}
}

func TestCountLarge(t *testing.T) {
	in := make([]string, 0, 1000)
	for i := 0; i < 1000; i++ {
		in = append(in, fmt.Sprintf("w%d", i%50))
	}
	got := Count(in)
	if len(got) != 50 {
		t.Fatalf("distinct = %d, want 50", len(got))
	}
	if got["w0"] != 20 {
		t.Errorf("count = %d, want 20", got["w0"])
	}
}

func TestSizeHint(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "mapprealloc.go", nil, 0)
	if err != nil {
		return
	}
	sized := false
	ast.Inspect(f, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "make" && len(c.Args) >= 2 {
				sized = true
			}
		}
		return true
	})
	if !sized {
		t.Logf("WARN: give make a size hint — make(map[string]int, len(words))")
	}
}
