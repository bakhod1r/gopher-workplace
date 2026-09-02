package parallelmap

import (
	"runtime"
	"sync/atomic"
	"testing"
)

type peakOp struct {
	running *int64
	peak    *int64
	spin    int
}

func (p peakOp) Apply(v int) int {
	cur := atomic.AddInt64(p.running, 1)
	for {
		old := atomic.LoadInt64(p.peak)
		if cur <= old || atomic.CompareAndSwapInt64(p.peak, old, cur) {
			break
		}
	}

	acc := v
	for i := 0; i < p.spin; i++ {
		acc = acc*31 + i
	}

	atomic.AddInt64(p.running, -1)
	return acc
}

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

func TestApply(t *testing.T) {
	if got := (SquareOp{}).Apply(3); got != 9 {
		t.Errorf("Apply = %d, want 9", got)
	}
}

func TestMapParallelBasic(t *testing.T) {
	if got := MapParallel(SquareOp{}, []int{1, 2, 3}); !eq(got, []int{1, 4, 9}) {
		t.Errorf("MapParallel = %v, want [1 4 9]", got)
	}
}

func TestMatchesSequential(t *testing.T) {
	vs := make([]int, 1000)
	for i := range vs {
		vs[i] = i
	}
	if got, want := MapParallel(SquareOp{}, vs), MapSeq(SquareOp{}, vs); !eq(got, want) {
		t.Error("MapParallel disagrees with MapSeq")
	}
}

func TestSmallerThanWorkerCount(t *testing.T) {
	if got := MapParallel(SquareOp{}, []int{5}); !eq(got, []int{25}) {
		t.Errorf("MapParallel = %v, want [25]", got)
	}
	if got := MapParallel(SquareOp{}, []int{2, 3}); !eq(got, []int{4, 9}) {
		t.Errorf("MapParallel = %v, want [4 9]", got)
	}
}

func TestEmpty(t *testing.T) {
	if got := MapParallel(SquareOp{}, nil); len(got) != 0 {
		t.Errorf("MapParallel = %v, want empty", got)
	}
}

func TestActuallyParallel(t *testing.T) {
	if runtime.GOMAXPROCS(0) < 2 {
		t.Skip("needs more than one P to observe parallelism")
	}

	var running, peak int64
	op := peakOp{running: &running, peak: &peak, spin: 20000}

	vs := make([]int, 4096)
	MapParallel(op, vs)

	if peak < 2 {
		t.Errorf("peak parallelism = %d; the workers never overlapped", peak)
	}
	if running != 0 {
		t.Errorf("%d workers still running after MapParallel returned", running)
	}
}

func TestEveryElementProcessedOnce(t *testing.T) {
	vs := make([]int, 997) // a prime length, so chunks do not divide evenly
	for i := range vs {
		vs[i] = i
	}
	got := MapParallel(SquareOp{}, vs)
	for i := range vs {
		if got[i] != i*i {
			t.Fatalf("got[%d] = %d, want %d", i, got[i], i*i)
		}
	}
}

func BenchmarkMapSeq(b *testing.B) {
	vs := make([]int, 100000)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MapSeq(SquareOp{}, vs)
	}
}

func BenchmarkMapParallel(b *testing.B) {
	vs := make([]int, 100000)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MapParallel(SquareOp{}, vs)
	}
}
