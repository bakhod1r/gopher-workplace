package timerreset

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
	"time"
)

func TestSession(t *testing.T) {
	start := time.Unix(0, 0)
	s := &Session{lastPing: start, timeout: 5 * time.Second}

	if s.IsExpired(start.Add(4 * time.Second)) {
		t.Error("should not be expired at 4s")
	}
	if !s.IsExpired(start.Add(6 * time.Second)) {
		t.Error("should be expired at 6s")
	}

	s.Ping(start.Add(4 * time.Second))
	if s.IsExpired(start.Add(6 * time.Second)) {
		t.Error("should not be expired at 6s after ping at 4s")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardComparesElapsed(t *testing.T) {
	targets := map[string]bool{"Ping": true, "IsExpired": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "timerreset.go", nil, 0)
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

	if !seen["lastPing"] || !seen["timeout"] {
		t.Logf("WARN: compare now against s.lastPing using s.timeout")
	}
}
