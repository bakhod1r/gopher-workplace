package lrumemcache

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"testing"
)

func TestLRU(t *testing.T) {
	l := New(2)
	l.Put("a", 1)
	l.Put("b", 2)

	if v, ok := l.Get("a"); !ok || v != 1 {
		t.Error("expected a=1")
	}

	l.Put("c", 3) // Evicts "b"

	if _, ok := l.Get("b"); ok {
		t.Error("expected b evicted")
	}
	if v, ok := l.Get("c"); !ok || v != 3 {
		t.Error("expected c=3")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardKeepsMapAndListInStep(t *testing.T) {
	targets := map[string]bool{"Get": true, "Put": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "lrumemcache.go", nil, 0)
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

	if !seen["cache"] || !seen["remove"] || !seen["insert"] {
		t.Logf("WARN: every operation must keep l.cache and the list in step (remove/insert, and delete on eviction)")
	}
}

// The cache is a fixed-memory structure: no matter how many distinct keys pass
// through it, it must never hold more than cap entries — in the map or in the
// list.
func TestMemoryStaysBounded(t *testing.T) {
	const cap = 100
	l := New(cap)

	for i := 0; i < 100000; i++ {
		l.Put(strconv.Itoa(i), i)
	}

	if len(l.cache) != cap {
		t.Errorf("map holds %d entries, want %d", len(l.cache), cap)
	}

	n := 0
	for node := l.head.next; node != l.tail; node = node.next {
		n++
		if n > cap {
			t.Fatalf("list holds more than %d nodes — an evicted node was not unlinked", cap)
		}
	}
	if n != cap {
		t.Errorf("list holds %d nodes, want %d", n, cap)
	}
}

// The most recently used keys must be the survivors.
func TestRecentKeysSurvive(t *testing.T) {
	l := New(3)
	for i := 0; i < 10; i++ {
		l.Put(strconv.Itoa(i), i)
	}

	for i := 7; i < 10; i++ {
		if v, ok := l.Get(strconv.Itoa(i)); !ok || v != i {
			t.Errorf("key %d = (%d, %v), want (%d, true)", i, v, ok, i)
		}
	}
	if _, ok := l.Get("6"); ok {
		t.Error("key 6 should have been evicted")
	}
}
