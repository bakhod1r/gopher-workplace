package middleware

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestStackThen(t *testing.T) {
	mw1 := func(next Handler) Handler {
		return func(req string) string { return "1(" + next(req) + ")1" }
	}
	mw2 := func(next Handler) Handler {
		return func(req string) string { return "2(" + next(req) + ")2" }
	}

	handler := Stack{mw1, mw2}.Then(func(req string) string { return "H:" + req })

	got := handler("req")
	want := "1(2(H:req)2)1"
	if got != want {
		t.Errorf("Then() = %q, want %q", got, want)
	}
}

func TestEmptyStackIsIdentity(t *testing.T) {
	base := func(req string) string { return "H:" + req }

	if got := (Stack{}).Then(base)("req"); got != "H:req" {
		t.Errorf("empty stack = %q, want %q", got, "H:req")
	}
}

func TestSingleMiddleware(t *testing.T) {
	mw := func(next Handler) Handler {
		return func(req string) string { return "1(" + next(req) + ")1" }
	}

	got := Stack{mw}.Then(func(req string) string { return "H:" + req })("req")
	if want := "1(H:req)1"; got != want {
		t.Errorf("Then() = %q, want %q", got, want)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardFoldsTheStack(t *testing.T) {
	targets := map[string]bool{"Then": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "middleware.go", nil, 0)
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

	if !seen["s"] || !seen["next"] {
		t.Logf("WARN: fold the stack over next; don't rebuild the chain by hand")
	}
}
