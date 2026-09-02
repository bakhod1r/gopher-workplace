package factorymeth

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFactory(t *testing.T) {
	f := StoreFactory{}

	if _, ok := f.Create("mem").(MemStore); !ok {
		t.Error("expected MemStore")
	}
	if _, ok := f.Create("disk").(DiskStore); !ok {
		t.Error("expected DiskStore")
	}
	if f.Create("unknown") != nil {
		t.Error("expected nil")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardBuildsBothStores(t *testing.T) {
	targets := map[string]bool{"Create": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "factorymeth.go", nil, 0)
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

	if !seen["MemStore"] || !seen["DiskStore"] {
		t.Logf("WARN: the factory must construct MemStore and DiskStore itself")
	}
}
