package pointerwalk

import (
	"testing"
	"unsafe"
)

var sink int64

func TestSumInt32(t *testing.T) {
	a := []int32{1, 2, 3, 4}
	if got := SumInt32(unsafe.SliceData(a), 3); got != 6 {
		t.Errorf("SumInt32 = %d, want 6", got)
	}
	if got := SumInt32(unsafe.SliceData(a), 4); got != 10 {
		t.Errorf("SumInt32 = %d, want 10", got)
	}
}

func TestSumInt32Edges(t *testing.T) {
	a := []int32{1, 2}
	if got := SumInt32(unsafe.SliceData(a), 0); got != 0 {
		t.Errorf("SumInt32 = %d, want 0", got)
	}
	if got := SumInt32(unsafe.SliceData(a), -1); got != 0 {
		t.Errorf("SumInt32 = %d, want 0", got)
	}
	if got := SumInt32(nil, 3); got != 0 {
		t.Errorf("SumInt32(nil) = %d, want 0", got)
	}
}

func TestSumInt32Negative(t *testing.T) {
	a := []int32{-5, 5, -1}
	if got := SumInt32(unsafe.SliceData(a), 3); got != -1 {
		t.Errorf("SumInt32 = %d, want -1", got)
	}
}

func TestSumInt32WideAccumulator(t *testing.T) {
	a := make([]int32, 8)
	for i := range a {
		a[i] = 1 << 30
	}
	if got := SumInt32(unsafe.SliceData(a), 8); got != 8<<30 {
		t.Errorf("SumInt32 = %d, want %d: the total must not overflow", got, int64(8)<<30)
	}
}

func TestSumInt32AllocatesNothing(t *testing.T) {
	a := make([]int32, 1024)
	for i := range a {
		a[i] = int32(i)
	}
	p := unsafe.SliceData(a)
	if n := testing.AllocsPerRun(100, func() { sink = SumInt32(p, 1024) }); n != 0 {
		t.Errorf("SumInt32 made %v allocations, want 0", n)
	}
}
