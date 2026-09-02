package actorrouter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestRouter(t *testing.T) {
	r := &Router{
		workers: []*Worker{
			{Inbox: make(chan int, 10)},
			{Inbox: make(chan int, 10)},
		},
	}

	r.Route(1)
	r.Route(2)
	r.Route(3)

	if <-r.workers[0].Inbox != 1 {
		t.Error("worker 0 missed 1")
	}
	if <-r.workers[1].Inbox != 2 {
		t.Error("worker 1 missed 2")
	}
	if <-r.workers[0].Inbox != 3 {
		t.Error("worker 0 missed 3")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardRoundRobins(t *testing.T) {
	targets := map[string]bool{"Route": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "actorrouter.go", nil, 0)
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

	if !seen["idx"] || !seen["workers"] {
		t.Logf("WARN: deliver to workers[idx] and advance idx modulo the pool size")
	}
}
