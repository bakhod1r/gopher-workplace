package ratelimit

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestLimiter(t *testing.T) {
	l := NewLimiter(2)

	if !l.Allow() {
		t.Error("expected true")
	}
	if !l.Allow() {
		t.Error("expected true")
	}
	if l.Allow() {
		t.Error("expected false (empty)")
	}

	l.Refill(1)
	if !l.Allow() {
		t.Error("expected true (refilled)")
	}
	if l.Allow() {
		t.Error("expected false (empty again)")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardGuardsTokens(t *testing.T) {
	targets := map[string]bool{"Allow": true, "Refill": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "ratelimit.go", nil, 0)
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

	if !seen["tokens"] || !seen["Lock"] {
		t.Logf("WARN: read and change l.tokens under the mutex")
	}
}
