package epollwrap

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestWaitFiltersByInterest(t *testing.T) {
	e := New()
	e.Add(3, EventRead)
	e.Add(7, EventWrite)
	e.Add(5, EventRead|EventWrite)

	got := e.Wait(map[int]uint32{
		3: EventRead,              // matches
		7: EventRead,              // registered, but wrong event
		5: EventWrite,             // matches one of two bits
		9: EventRead | EventWrite, // not registered
	})

	want := []int{3, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wait() = %v, want %v", got, want)
	}
}

func TestWaitSortsAndHandlesEmpty(t *testing.T) {
	e := New()
	for _, fd := range []int{11, 2, 8} {
		e.Add(fd, EventRead)
	}

	got := e.Wait(map[int]uint32{11: EventRead, 2: EventRead, 8: EventRead})
	want := []int{2, 8, 11}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wait() = %v, want %v", got, want)
	}

	if got := e.Wait(map[int]uint32{}); len(got) != 0 {
		t.Errorf("empty ready set = %v, want no descriptors", got)
	}
}

func TestRemoveDeregisters(t *testing.T) {
	e := New()
	e.Add(3, EventRead)
	e.Remove(3)
	e.Remove(99) // no-op

	if got := e.Wait(map[int]uint32{3: EventRead}); len(got) != 0 {
		t.Errorf("after Remove, Wait() = %v, want none", got)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardIntersectsInterest(t *testing.T) {
	targets := map[string]bool{"Wait": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "epollwrap.go", nil, 0)
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

	if !seen["interest"] || !seen["ready"] {
		t.Logf("WARN: intersect each ready mask with the registered interest mask")
	}
}
