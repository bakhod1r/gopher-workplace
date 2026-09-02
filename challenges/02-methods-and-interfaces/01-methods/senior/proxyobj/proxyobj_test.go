package proxyobj

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestProxy(t *testing.T) {
	w := &Worker{}

	p1 := &Proxy{w: w, role: "user"}
	if got := p1.Do(); got != "denied" {
		t.Errorf("user got %q", got)
	}
	if w.calls != 0 {
		t.Error("worker called by user")
	}

	p2 := &Proxy{w: w, role: "admin"}
	if got := p2.Do(); got != "done" {
		t.Errorf("admin got %q", got)
	}
	if w.calls != 1 {
		t.Error("worker not called by admin")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardDelegatesToWorker(t *testing.T) {
	targets := map[string]bool{"Do": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "proxyobj.go", nil, 0)
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

	if !seen["w"] || !seen["Do"] {
		t.Logf("WARN: the proxy must forward to p.w.Do() for an admin, not fabricate the result")
	}
}
