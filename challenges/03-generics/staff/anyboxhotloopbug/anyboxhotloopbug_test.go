package anyboxhotloopbug

import (
	"testing"
	"time"
)

var sink int

func TestSumBasic(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum = %d, want 6", got)
	}
}

func TestSumFloatAndEmpty(t *testing.T) {
	if got := Sum([]float64{1.5, 2.5}); got != 4 {
		t.Errorf("Sum = %v, want 4", got)
	}
	if got := Sum([]int{}); got != 0 {
		t.Errorf("Sum = %d, want 0", got)
	}
}

func TestSumDoesNotAllocatePerElement(t *testing.T) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = 1000 + i
	}
	allocs := testing.AllocsPerRun(5, func() { sink = Sum(data) })
	if allocs > 2 {
		t.Errorf("Sum allocated %.0f times per call, want at most 2", allocs)
	}
}

func TestSumScale(t *testing.T) {
	const n = 3_000_000
	data := make([]int, n)
	for i := range data {
		data[i] = i
	}
	start := time.Now()
	got := Sum(data)
	elapsed := time.Since(start)
	want := n * (n - 1) / 2
	if got != want {
		t.Fatalf("Sum = %d, want %d", got, want)
	}
	if elapsed > 200*time.Millisecond {
		t.Errorf("Sum of %d elements took %v, want under 200ms", n, elapsed)
	}
}
