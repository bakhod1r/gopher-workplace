package iterchain

import "testing"

func eq(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range want {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func double(v int) int { return v * 2 }

func even(v int) bool { return v%2 == 0 }

func TestMapIter(t *testing.T) {
	src := &SliceIter{Data: []int{1, 2, 3}}
	got := Collect(&MapIter{Inner: src, Fn: double})
	if !eq(got, []int{2, 4, 6}) {
		t.Errorf("got %v, want [2 4 6]", got)
	}
}

func TestFilterIter(t *testing.T) {
	src := &SliceIter{Data: []int{1, 2, 3, 4}}
	got := Collect(&FilterIter{Inner: src, Pred: even})
	if !eq(got, []int{2, 4}) {
		t.Errorf("got %v, want [2 4]", got)
	}
}

func TestChain(t *testing.T) {
	src := &SliceIter{Data: []int{1, 2, 3}}
	chain := &FilterIter{
		Inner: &MapIter{Inner: src, Fn: double},
		Pred:  even,
	}
	if got := Collect(chain); !eq(got, []int{2, 4, 6}) {
		t.Errorf("got %v, want [2 4 6]", got)
	}
}

func TestFilterEverything(t *testing.T) {
	src := &SliceIter{Data: []int{1, 3, 5}}
	if got := Collect(&FilterIter{Inner: src, Pred: even}); len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
	if src.Reads != 3 {
		t.Errorf("Reads = %d, want 3", src.Reads)
	}
}

func TestLazyUntilCollected(t *testing.T) {
	src := &SliceIter{Data: []int{1, 2, 3}}
	_ = &FilterIter{Inner: &MapIter{Inner: src, Fn: double}, Pred: even}

	if src.Reads != 0 {
		t.Errorf("building the chain read %d elements, want 0", src.Reads)
	}
}

func TestEachElementReadOnce(t *testing.T) {
	src := &SliceIter{Data: []int{1, 2, 3, 4, 5}}
	chain := &MapIter{
		Inner: &FilterIter{Inner: &MapIter{Inner: src, Fn: double}, Pred: even},
		Fn:    double,
	}
	Collect(chain)

	if src.Reads != 5 {
		t.Errorf("Reads = %d, want 5 (each element pulled once)", src.Reads)
	}
}

func TestEmptySource(t *testing.T) {
	src := &SliceIter{}
	if got := Collect(&MapIter{Inner: src, Fn: double}); len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func BenchmarkChain(b *testing.B) {
	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		src := &SliceIter{Data: data}
		Collect(&FilterIter{Inner: &MapIter{Inner: src, Fn: double}, Pred: even})
	}
}
