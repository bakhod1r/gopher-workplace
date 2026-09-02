package shardedmap

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"sync"
	"testing"
)

func TestSetGet(t *testing.T) {
	m := New(8)
	m.Set("a", 1)
	m.Set("b", 2)

	if v, ok := m.Get("a"); !ok || v != 1 {
		t.Errorf("Get(a) = (%d, %v), want (1, true)", v, ok)
	}
	if v, ok := m.Get("missing"); ok || v != 0 {
		t.Errorf("Get(missing) = (%d, %v), want (0, false)", v, ok)
	}
	if m.Len() != 2 {
		t.Errorf("Len() = %d, want 2", m.Len())
	}
}

func TestConcurrentDistinctKeys(t *testing.T) {
	m := New(32)
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set(strconv.Itoa(i), i)
		}(i)
	}
	wg.Wait()

	if m.Len() != 1000 {
		t.Errorf("Len() = %d, want 1000", m.Len())
	}
	for i := 0; i < 1000; i++ {
		if v, ok := m.Get(strconv.Itoa(i)); !ok || v != i {
			t.Fatalf("Get(%d) = (%d, %v), want (%d, true)", i, v, ok, i)
		}
	}
}

func TestConcurrentSameKey(t *testing.T) {
	m := New(32)
	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set("hot", i)
			m.Get("hot")
		}(i)
	}
	wg.Wait()

	if _, ok := m.Get("hot"); !ok {
		t.Error("expected key hot to exist")
	}
}

// Warning guard (GENERATION.md §5f): advisory only — it never fails a Run, but
// an unresolved WARN blocks Submit. The file is parsed with mode 0 so comments
// are excluded: naming the rule in a comment cannot satisfy the check. Only the
// bodies of the functions under task are inspected, so identifiers that already
// appear elsewhere in the file do not count.
func TestGuardLocksTheShard(t *testing.T) {
	targets := map[string]bool{"Set": true, "Get": true}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "shardedmap.go", nil, 0)
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

	if !seen["getShard"] || !seen["RLock"] {
		t.Logf("WARN: route through getShard and take that shard's lock - Get should use RLock")
	}
}

// Sharding must hold up under sustained mixed read/write load from many
// goroutines — correct results, and clean under -race.
func TestMixedLoad(t *testing.T) {
	const (
		writers      = 8
		perGoroutine = 5000
	)

	m := New(32)
	var wg sync.WaitGroup

	for w := 0; w < writers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			for i := 0; i < perGoroutine; i++ {
				key := strconv.Itoa(w*perGoroutine + i)
				m.Set(key, i)
				if v, ok := m.Get(key); !ok || v != i {
					t.Errorf("Get(%s) = (%d, %v), want (%d, true)", key, v, ok, i)
					return
				}
			}
		}(w)
	}
	wg.Wait()

	if got := m.Len(); got != writers*perGoroutine {
		t.Errorf("Len() = %d, want %d", got, writers*perGoroutine)
	}
}

// Keys must spread across shards — a router that sends everything to one shard
// is correct but defeats the entire design.
func TestKeysSpreadAcrossShards(t *testing.T) {
	m := New(16)
	for i := 0; i < 10000; i++ {
		m.Set(strconv.Itoa(i), i)
	}

	used := 0
	for _, s := range m.shards {
		if len(s.data) > 0 {
			used++
		}
	}
	if used < len(m.shards) {
		t.Errorf("%d of %d shards used — keys are not spreading", used, len(m.shards))
	}
}
