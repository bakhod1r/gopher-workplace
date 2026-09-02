package chainresp

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestChain(t *testing.T) {
	h10 := &H10{}
	h20 := &H20{}
	h10.SetNext(h20)

	if got := h10.Handle(10); got != "ten" {
		t.Errorf("10 = %q", got)
	}
	if got := h10.Handle(20); got != "twenty" {
		t.Errorf("20 = %q", got)
	}
	if got := h10.Handle(30); got != "unhandled" {
		t.Errorf("30 = %q", got)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardForwardsDownTheChain(t *testing.T) {
	targets := map[string]bool{"Handle": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "chainresp.go", nil, 0)
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

	if !seen["Next"] {
		t.Logf("WARN: forward with h.Next(req) instead of producing the fallback yourself")
	}
}
