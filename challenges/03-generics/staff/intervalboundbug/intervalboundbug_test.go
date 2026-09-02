package intervalboundbug

import (
	"testing"
	"time"
)

func TestAdjacentRanges(t *testing.T) {
	var a Accum[int]
	a.Add(0, 3, 1)
	a.Add(3, 6, 1)
	for x := 0; x < 6; x++ {
		if got := a.At(x); got != 1 {
			t.Errorf("At(%d) = %d, want 1", x, got)
		}
	}
	if got := a.At(6); got != 0 {
		t.Errorf("At(6) = %d, want 0", got)
	}
	if got := a.Total(); got != 6 {
		t.Errorf("Total = %d, want 6", got)
	}
}

func TestEmptyRange(t *testing.T) {
	var a Accum[int]
	a.Add(4, 4, 5)
	if got := a.At(4); got != 0 {
		t.Errorf("At(4) = %d, want 0", got)
	}
	if got := a.Touched(); got != 0 {
		t.Errorf("Touched = %d, want 0", got)
	}
}

func TestTilingScale(t *testing.T) {
	const ranges = 20000
	const width = 10
	start := time.Now()
	var a Accum[int]
	for i := 0; i < ranges; i++ {
		a.Add(i*width, i*width+width, 1)
	}
	if got := a.Touched(); got != ranges*width {
		t.Fatalf("Touched = %d, want %d", got, ranges*width)
	}
	if got := a.Total(); got != ranges*width {
		t.Fatalf("Total = %d, want %d", got, ranges*width)
	}
	for _, x := range []int{0, width, 2*width - 1, ranges*width - 1} {
		if got := a.At(x); got != 1 {
			t.Fatalf("At(%d) = %d, want 1", x, got)
		}
	}
	if d := time.Since(start); d > 3*time.Second {
		t.Fatalf("20k tiled ranges took %v, want under 3s", d)
	}
}
