package chunkworkers

import (
	"runtime"
	"testing"
)

func TestSumParallel(t *testing.T) {
	if got := SumParallel([]int{1, 2, 3, 4}, 2); got != 10 {
		t.Errorf("SumParallel = %d, want 10", got)
	}
	if got := SumParallel(nil, 4); got != 0 {
		t.Errorf("SumParallel = %d, want 0", got)
	}
	if got := SumParallel([]int{5}, 8); got != 5 {
		t.Errorf("SumParallel = %d, want 5: more workers than elements", got)
	}
	if got := SumParallel([]int{1, 2, 3}, 0); got != 6 {
		t.Errorf("SumParallel = %d, want 6: workers < 1 must still work", got)
	}
}

func TestSumParallelMatchesSerial(t *testing.T) {
	s := make([]int, 100003)
	var want int64
	for i := range s {
		s[i] = i % 977
		want += int64(s[i])
	}
	for _, w := range []int{1, 2, 3, 7, 16} {
		if got := SumParallel(s, w); got != want {
			t.Fatalf("SumParallel(_, %d) = %d, want %d: the chunks do not cover the input exactly once", w, got, want)
		}
	}
}

func TestSumParallelDoesNotCopyTheInput(t *testing.T) {
	s := make([]int, 1<<20)
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	SumParallel(s, 8)
	runtime.ReadMemStats(&after)
	if used := after.TotalAlloc - before.TotalAlloc; used > 1<<16 {
		t.Errorf("allocated %d bytes for an 8 MiB input, want under 64 KiB: pass views, not copies", used)
	}
}
