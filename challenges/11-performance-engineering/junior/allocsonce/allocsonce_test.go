package allocsonce

import "testing"

var sinkI int
var sinkB bool

func TestLookup(t *testing.T) {
	ix := &Index{Words: []string{"a", "b", "c"}}
	if got, ok := ix.Lookup("b"); got != 1 || !ok {
		t.Errorf("Lookup(b) = %d, %v, want 1, true", got, ok)
	}
	if got, ok := ix.Lookup("zz"); got != 0 || ok {
		t.Errorf("Lookup(zz) = %d, %v, want 0, false", got, ok)
	}
}

func TestLookupFirstOccurrenceWins(t *testing.T) {
	ix := &Index{Words: []string{"a", "b", "a"}}
	if got, _ := ix.Lookup("a"); got != 0 {
		t.Errorf("Lookup(a) = %d, want 0", got)
	}
}

func TestLookupEmptyIndex(t *testing.T) {
	ix := &Index{}
	if _, ok := ix.Lookup("a"); ok {
		t.Error("Lookup on an empty index reported found")
	}
}

func TestLookupIsAllocationFreeAfterTheFirstCall(t *testing.T) {
	words := make([]string, 0, 500)
	for i := 0; i < 500; i++ {
		words = append(words, string(rune('a'+i%26))+string(rune('a'+i/26)))
	}
	ix := &Index{Words: words}
	ix.Lookup("aa") // build the index
	allocs := testing.AllocsPerRun(100, func() { sinkI, sinkB = ix.Lookup("ba") })
	if allocs != 0 {
		t.Errorf("warm Lookup made %v allocations, want 0 — the map must be built once and reused", allocs)
	}
}
