package boxvalues

import "testing"

var sink int64

func TestTotal(t *testing.T) {
	if got := Total([]int{1, 2, 3}); got != 6 {
		t.Errorf("Total = %d, want 6", got)
	}
	if got := Total(nil); got != 0 {
		t.Errorf("Total = %d, want 0", got)
	}
	if got := Total([]int{-5, 5}); got != 0 {
		t.Errorf("Total = %d, want 0", got)
	}
}

func TestTotalWideAccumulator(t *testing.T) {
	vals := make([]int, 8)
	for i := range vals {
		vals[i] = 1 << 40
	}
	if got := Total(vals); got != 8<<40 {
		t.Errorf("Total = %d, want %d", got, int64(8)<<40)
	}
}

func TestTotalAllocatesNothing(t *testing.T) {
	vals := make([]int, 64)
	for i := range vals {
		vals[i] = 1000 + i
	}
	if n := testing.AllocsPerRun(100, func() { sink = Total(vals) }); n != 0 {
		t.Errorf("Total made %v allocations, want 0: the values are being boxed", n)
	}
}

func BenchmarkTotal(b *testing.B) {
	vals := make([]int, 4096)
	for i := range vals {
		vals[i] = i
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = Total(vals)
	}
}
