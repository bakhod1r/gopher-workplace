package movingavg

import (
	"math"
	"reflect"
	"testing"
)

func nearAll(got, want []float64) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if math.Abs(got[i]-want[i]) > 1e-9 {
			return false
		}
	}
	return true
}

func TestWindow(t *testing.T) {
	if got := Window([]float64{1, 2, 3, 4}, 2); !nearAll(got, []float64{1.5, 2.5, 3.5}) {
		t.Errorf("Window = %v, want [1.5 2.5 3.5]", got)
	}
	if got := Window([]float64{1, 2, 3, 4}, 4); !nearAll(got, []float64{2.5}) {
		t.Errorf("Window = %v, want [2.5]", got)
	}
	if got := Window([]float64{5, 5, 5}, 1); !nearAll(got, []float64{5, 5, 5}) {
		t.Errorf("Window = %v, want [5 5 5]", got)
	}
}

func TestWindowGuards(t *testing.T) {
	for _, c := range []struct {
		samples []float64
		n       int
	}{{[]float64{1, 2}, 0}, {[]float64{1, 2}, -1}, {[]float64{1, 2}, 3}, {nil, 2}} {
		got := Window(c.samples, c.n)
		if got == nil || len(got) != 0 {
			t.Errorf("Window(%v, %d) = %v, want empty non-nil slice", c.samples, c.n, got)
		}
	}
}

func TestWindowDoesNotModifyInput(t *testing.T) {
	in := []float64{1, 2, 3, 4}
	before := append([]float64(nil), in...)
	Window(in, 2)
	if !reflect.DeepEqual(in, before) {
		t.Errorf("input changed: %v, want %v", in, before)
	}
}

func TestWindowIsNumericallySane(t *testing.T) {
	samples := make([]float64, 10_000)
	for i := range samples {
		samples[i] = float64(i % 7)
	}
	got := Window(samples, 100)
	if len(got) != 10_000-100+1 {
		t.Fatalf("len = %d, want %d", len(got), 10_000-100+1)
	}
	// Every window of 100 samples over a period-7 pattern has a mean near 3.
	for i, v := range got {
		if v < 2.5 || v > 3.5 {
			t.Fatalf("window %d mean = %v, out of the expected range", i, v)
		}
	}
}

func TestSmoothest(t *testing.T) {
	if got := Smoothest([]float64{3, 1, 1, 5}); got != 1 {
		t.Errorf("Smoothest = %d, want 1", got)
	}
	if got := Smoothest(nil); got != -1 {
		t.Errorf("Smoothest(nil) = %d, want -1", got)
	}
	if got := Smoothest([]float64{9}); got != 0 {
		t.Errorf("Smoothest = %d, want 0", got)
	}
}
