package warmupdrop

import (
	"math"
	"reflect"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestDrop(t *testing.T) {
	if got := Drop([]float64{9, 1, 1}, 1); !reflect.DeepEqual(got, []float64{1, 1}) {
		t.Errorf("Drop = %v, want [1 1]", got)
	}
	if got := Drop([]float64{1, 2}, 0); !reflect.DeepEqual(got, []float64{1, 2}) {
		t.Errorf("Drop = %v, want [1 2]", got)
	}
	if got := Drop([]float64{1, 2}, -3); !reflect.DeepEqual(got, []float64{1, 2}) {
		t.Errorf("Drop = %v, want [1 2]", got)
	}
}

func TestDropBeyondLength(t *testing.T) {
	for _, n := range []int{2, 5} {
		got := Drop([]float64{1, 2}, n)
		if got == nil || len(got) != 0 {
			t.Errorf("Drop(_, %d) = %v, want empty non-nil slice", n, got)
		}
	}
}

func TestDropDoesNotModifyInput(t *testing.T) {
	in := []float64{9, 1, 1}
	before := append([]float64(nil), in...)
	got := Drop(in, 1)
	got[0] = 99
	if !reflect.DeepEqual(in, before) {
		t.Errorf("input changed: %v, want %v", in, before)
	}
}

func TestStableMean(t *testing.T) {
	if got := StableMean([]float64{100, 2, 4}, 1); !near(got, 3) {
		t.Errorf("StableMean = %v, want 3", got)
	}
	if got := StableMean([]float64{100, 2, 4}, 0); !near(got, 106.0/3) {
		t.Errorf("StableMean = %v, want %v", got, 106.0/3)
	}
	if got := StableMean([]float64{1, 2}, 5); got != 0 {
		t.Errorf("StableMean with everything dropped = %v, want 0", got)
	}
}
