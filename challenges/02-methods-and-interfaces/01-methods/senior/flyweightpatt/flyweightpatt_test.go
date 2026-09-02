package flyweightpatt

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFlyweight(t *testing.T) {
	f := &FlyweightFactory{fonts: make(map[string]*FontData)}

	f1 := f.Get("Arial")
	f2 := f.Get("Arial")
	f3 := f.Get("Times")

	if f1 != f2 {
		t.Error("expected same Arial instance")
	}
	if f1 == f3 {
		t.Error("expected different Times instance")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardSharesInstances(t *testing.T) {
	targets := map[string]bool{"Get": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "flyweightpatt.go", nil, 0)
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

	if !seen["fonts"] {
		t.Logf("WARN: look the name up in f.fonts and store the new instance there - sharing is the point")
	}
}
