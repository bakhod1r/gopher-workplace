package hotpathgeneric

import "testing"

var sink int64

type myInt int

func TestTotal(t *testing.T) {
	if got := Total([]int{1, 2, 3}); got != 6 {
		t.Errorf("Total = %d, want 6", got)
	}
	if got := Total([]int64{1 << 40, 1 << 40}); got != 1<<41 {
		t.Errorf("Total = %d, want %d", got, int64(1)<<41)
	}
	if got := Total([]int32{-5, 5}); got != 0 {
		t.Errorf("Total = %d, want 0", got)
	}
	if got := Total[myInt](nil); got != 0 {
		t.Errorf("Total = %d, want 0", got)
	}
}

func TestTotalAcceptsNamedTypes(t *testing.T) {
	if got := Total([]myInt{2, 3}); got != 5 {
		t.Errorf("Total = %d, want 5", got)
	}
}

func TestTotalAllocatesNothing(t *testing.T) {
	vals := make([]int, 1024)
	for i := range vals {
		vals[i] = i
	}
	if n := testing.AllocsPerRun(100, func() { sink = Total(vals) }); n != 0 {
		t.Errorf("Total made %v allocations, want 0", n)
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
