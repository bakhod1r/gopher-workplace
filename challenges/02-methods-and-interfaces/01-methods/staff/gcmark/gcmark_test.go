package gcmark

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestGCMark(t *testing.T) {
	o1 := &Object{}
	o2 := &Object{}
	o1.Refs = append(o1.Refs, o2)
	o1.Mark()
	if !o2.Marked {
		t.Error("expected o2 marked")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardTracesRefs(t *testing.T) {
	targets := map[string]bool{"Mark": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "gcmark.go", nil, 0)
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

	if !seen["Refs"] || !seen["Mark"] {
		t.Logf("WARN: recurse into o.Refs with Mark(); the mark bit is also the visited set")
	}
}

// Object graphs have cycles. Marking must terminate on one — the mark bit is
// the visited set.
func TestCycleTerminates(t *testing.T) {
	a := &Object{}
	b := &Object{}
	a.Refs = append(a.Refs, b)
	b.Refs = append(b.Refs, a)

	a.Mark() // must not recurse forever

	if !a.Marked || !b.Marked {
		t.Errorf("marked a=%v b=%v, want both true", a.Marked, b.Marked)
	}
}

// A long chain must be traced completely without blowing up.
func TestDeepGraph(t *testing.T) {
	const n = 100000

	objs := make([]*Object, n)
	for i := range objs {
		objs[i] = &Object{}
	}
	for i := 0; i < n-1; i++ {
		objs[i].Refs = append(objs[i].Refs, objs[i+1])
	}
	// close the loop so the trace also has to survive a cycle at depth
	objs[n-1].Refs = append(objs[n-1].Refs, objs[0])

	objs[0].Mark()

	for i, o := range objs {
		if !o.Marked {
			t.Fatalf("object %d was not marked", i)
		}
	}
}

// Unreachable objects must stay unmarked — that is what makes them collectable.
func TestUnreachableStaysUnmarked(t *testing.T) {
	root := &Object{}
	garbage := &Object{}

	root.Mark()

	if garbage.Marked {
		t.Error("an unreferenced object was marked")
	}
}
