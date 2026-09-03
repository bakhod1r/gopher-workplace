package sharedview

import (
	"testing"
	"unsafe"
)

func makeInt64Bytes(vals []int64) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(vals))), len(vals)*8)
}

func TestSumParallelSmall(t *testing.T) {
	vals := []int64{1, 2, 3, 4}
	got, ok := SumParallel(makeInt64Bytes(vals), 2)
	if !ok || got != 10 {
		t.Errorf("SumParallel = %d, %v, want 10, true", got, ok)
	}
}

func TestSumParallelWorkerCounts(t *testing.T) {
	vals := make([]int64, 1001)
	var want int64
	for i := range vals {
		vals[i] = int64(i)
		want += int64(i)
	}
	b := makeInt64Bytes(vals)
	for _, w := range []int{0, 1, 2, 7, 64, 100000} {
		got, ok := SumParallel(b, w)
		if !ok || got != want {
			t.Fatalf("workers=%d: SumParallel = %d, %v, want %d, true", w, got, ok, want)
		}
	}
}

func TestSumParallelRejectsBadShapes(t *testing.T) {
	vals := []int64{1, 2}
	b := makeInt64Bytes(vals)
	for _, c := range []struct {
		name string
		in   []byte
	}{
		{"nil", nil},
		{"empty", b[:0]},
		{"length not a multiple of 8", b[:12]},
		{"misaligned", b[1:9]},
	} {
		if _, ok := SumParallel(c.in, 2); ok {
			t.Errorf("%s: reported ok, want false", c.name)
		}
	}
}

func TestSumParallelIsRepeatable(t *testing.T) {
	vals := make([]int64, 4096)
	for i := range vals {
		vals[i] = int64(i % 97)
	}
	b := makeInt64Bytes(vals)
	first, ok := SumParallel(b, 8)
	if !ok {
		t.Fatal("SumParallel reported false")
	}
	for i := 0; i < 20; i++ {
		got, _ := SumParallel(b, 8)
		if got != first {
			t.Fatalf("run %d = %d, want %d: the workers overlap", i, got, first)
		}
	}
}

func TestSumParallelDoesNotCopy(t *testing.T) {
	vals := make([]int64, 1<<16)
	b := makeInt64Bytes(vals)
	var sink int64
	n := testing.AllocsPerRun(20, func() { sink, _ = SumParallel(b, 4) })
	_ = sink
	if n > 12 {
		t.Errorf("SumParallel made %v allocations for a 512 KiB buffer, want a handful: view, do not copy", n)
	}
}
