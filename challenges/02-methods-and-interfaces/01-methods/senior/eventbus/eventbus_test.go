package eventbus

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"testing"
)

func TestEventBus(t *testing.T) {
	b := New()
	var got1, got2 string

	b.On("user.login", func(data string) { got1 = "A:" + data })
	b.On("user.login", func(data string) { got2 = "B:" + data })

	b.Emit("user.login", "alice")

	if got1 != "A:alice" || got2 != "B:alice" {
		t.Errorf("Emit failed: got1=%q got2=%q", got1, got2)
	}

	// Emit with no listeners should not panic
	b.Emit("user.logout", "bob")
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardStoresAndCallsListeners(t *testing.T) {
	targets := map[string]bool{"On": true, "Emit": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "eventbus.go", nil, 0)
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

	if !seen["listeners"] || !seen["append"] {
		t.Logf("WARN: On must append to b.listeners; Emit must call what it finds there")
	}
}

// Fan-out must reach every listener, and only the listeners for that event.
func TestLargeFanOut(t *testing.T) {
	const listeners = 10000

	b := New()
	hits := 0
	for i := 0; i < listeners; i++ {
		b.On("tick", func(string) { hits++ })
	}
	for i := 0; i < 100; i++ {
		b.On("other-"+strconv.Itoa(i), func(string) { t.Error("wrong event delivered") })
	}

	b.Emit("tick", "now")

	if hits != listeners {
		t.Errorf("%d listeners fired, want %d", hits, listeners)
	}
}

// Emitting to an event nobody subscribed to must be free — no map write, no
// allocation, no panic.
func TestEmitToUnknownEventIsAllocationFree(t *testing.T) {
	b := New()
	b.On("known", func(string) {})

	allocs := testing.AllocsPerRun(1000, func() {
		b.Emit("unknown", "x")
	})
	if allocs != 0 {
		t.Errorf("Emit to an unknown event allocated %.0f times per run, want 0", allocs)
	}
}
