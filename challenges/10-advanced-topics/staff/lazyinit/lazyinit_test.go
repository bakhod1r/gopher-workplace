package lazyinit

import (
	"strconv"
	"sync"
	"testing"
)

func pairs(n int) [][2]string {
	out := make([][2]string, n)
	for i := range out {
		out[i] = [2]string{"k" + strconv.Itoa(i), "v"}
	}
	return out
}

func TestLookup(t *testing.T) {
	tbl := NewTable([][2]string{{"a", "1"}, {"b", "2"}})
	if i, ok := tbl.Lookup("a"); !ok || i != 0 {
		t.Errorf("Lookup(a) = %d, %v, want 0, true", i, ok)
	}
	if i, ok := tbl.Lookup("b"); !ok || i != 1 {
		t.Errorf("Lookup(b) = %d, %v, want 1, true", i, ok)
	}
	if _, ok := tbl.Lookup("missing"); ok {
		t.Error("Lookup(missing) reported ok, want false")
	}
}

func TestLookupEmptyTable(t *testing.T) {
	tbl := NewTable(nil)
	if _, ok := tbl.Lookup("a"); ok {
		t.Error("Lookup on an empty table reported ok, want false")
	}
}

func TestBuildsOncePerTable(t *testing.T) {
	tbl := NewTable(pairs(100))
	before := Builds.Load()
	for i := 0; i < 500; i++ {
		tbl.Lookup("k1")
	}
	if got := Builds.Load() - before; got != 1 {
		t.Errorf("the index was built %d times, want 1", got)
	}
}

func TestSeparateTablesBuildSeparately(t *testing.T) {
	before := Builds.Load()
	a := NewTable(pairs(4))
	b := NewTable(pairs(4))
	a.Lookup("k0")
	b.Lookup("k0")
	if got := Builds.Load() - before; got != 2 {
		t.Errorf("built %d times, want 2: the once must be per table", got)
	}
}

func TestConcurrentFirstUse(t *testing.T) {
	tbl := NewTable(pairs(1000))
	before := Builds.Load()
	var wg sync.WaitGroup
	const workers = 32
	wg.Add(workers)
	results := make([]bool, workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			_, ok := tbl.Lookup("k" + strconv.Itoa(w))
			results[w] = ok
		}(w)
	}
	wg.Wait()
	if got := Builds.Load() - before; got != 1 {
		t.Errorf("the index was built %d times under 32 concurrent first uses, want 1", got)
	}
	for w, ok := range results {
		if !ok {
			t.Fatalf("worker %d did not find its key: it read the index before it was built", w)
		}
	}
}
