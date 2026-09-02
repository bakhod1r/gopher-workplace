package memorize

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"testing"
)

func TestMemoizer(t *testing.T) {
	calls := 0
	fn := func(k string) string {
		calls++
		return k + "-val"
	}

	m := New(fn)

	if got := m.Get("a"); got != "a-val" || calls != 1 {
		t.Errorf("first call failed: %q, %d", got, calls)
	}

	if got := m.Get("a"); got != "a-val" || calls != 1 {
		t.Errorf("second call should be cached: %q, %d", got, calls)
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardCallsWrappedFunc(t *testing.T) {
	targets := map[string]bool{"Get": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "memorize.go", nil, 0)
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

	if !seen["fn"] || !seen["cache"] {
		t.Logf("WARN: a memoizer must consult m.cache and, on a miss, call m.fn")
	}
}

// The whole point of a memoizer is that the expensive function runs once per
// distinct key, no matter how often it is asked for.
func TestExpensiveFunctionRunsOncePerKey(t *testing.T) {
	const (
		keys    = 10000
		repeats = 10
	)

	calls := 0
	m := New(func(k string) string {
		calls++
		return k + "!"
	})

	for r := 0; r < repeats; r++ {
		for i := 0; i < keys; i++ {
			key := strconv.Itoa(i)
			if got, want := m.Get(key), key+"!"; got != want {
				t.Fatalf("Get(%s) = %q, want %q", key, got, want)
			}
		}
	}

	if calls != keys {
		t.Errorf("underlying function called %d times, want %d", calls, keys)
	}
}

// A cache hit is a map read — it must not re-enter the wrapped function or
// build anything.
func TestCacheHitIsAllocationFree(t *testing.T) {
	m := New(func(k string) string { return k + "!" })
	m.Get("warm")

	allocs := testing.AllocsPerRun(1000, func() {
		m.Get("warm")
	})
	if allocs != 0 {
		t.Errorf("cache hit allocated %.0f times per run, want 0", allocs)
	}
}
