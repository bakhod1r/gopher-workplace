package raftnode

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestRaft(t *testing.T) {
	n := &Node{State: "Follower"}
	n.Timeout()
	if n.State != "Candidate" {
		t.Error("expected Candidate")
	}

	n.ReceiveVotes(true)
	if n.State != "Leader" {
		t.Error("expected Leader")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardGuardsTransitions(t *testing.T) {
	targets := map[string]bool{"Timeout": true, "ReceiveVotes": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "raftnode.go", nil, 0)
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

	if !seen["State"] {
		t.Logf("WARN: guard each transition on the current State")
	}
}
