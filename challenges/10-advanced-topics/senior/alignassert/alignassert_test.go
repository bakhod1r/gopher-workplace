package alignassert

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
	"unsafe"
)

func TestCheck(t *testing.T) {
	if !Check() {
		t.Error("Check = false, want true: Value is at offset 0 of a 64-bit-aligned struct")
	}
}

func TestCheckIsRepeatable(t *testing.T) {
	for i := 0; i < 100; i++ {
		if !Check() {
			t.Fatalf("run %d: Check = false, want true", i)
		}
	}
}

func TestFixtureIsUnchanged(t *testing.T) {
	var c Counter
	if unsafe.Offsetof(c.Value) != 0 {
		t.Error("Value must stay the first field")
	}
	if unsafe.Sizeof(c.Value) != 8 {
		t.Error("Value must stay an int64")
	}
}

func TestCheckDerivesTheRequirement(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "alignassert.go", nil, 0)
	if err != nil {
		t.Skipf("cannot parse the source: %v", err)
	}
	usesAlignof := false
	ast.Inspect(f, func(n ast.Node) bool {
		if sel, ok := n.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok && id.Name == "unsafe" && sel.Sel.Name == "Alignof" {
				usesAlignof = true
			}
		}
		return true
	})
	if !usesAlignof {
		t.Error("the requirement must come from unsafe.Alignof, not from a hard-coded number")
	}
}
