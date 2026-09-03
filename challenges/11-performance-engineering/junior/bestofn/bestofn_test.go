package bestofn

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestBest(t *testing.T) {
	v, i := Best([]float64{5, 2, 9})
	if !near(v, 2) || i != 1 {
		t.Errorf("Best = %v, %d, want 2, 1", v, i)
	}
	v, i = Best([]float64{3})
	if !near(v, 3) || i != 0 {
		t.Errorf("Best = %v, %d, want 3, 0", v, i)
	}
}

func TestBestTieTakesTheEarliest(t *testing.T) {
	if _, i := Best([]float64{4, 1, 1, 1}); i != 1 {
		t.Errorf("index = %d, want 1", i)
	}
}

func TestBestEmpty(t *testing.T) {
	v, i := Best(nil)
	if v != 0 || i != -1 {
		t.Errorf("Best(nil) = %v, %d, want 0, -1", v, i)
	}
}

func TestSpread(t *testing.T) {
	if got := Spread([]float64{10, 20}); !near(got, 2) {
		t.Errorf("Spread = %v, want 2", got)
	}
	if got := Spread([]float64{7, 7, 7}); !near(got, 1) {
		t.Errorf("Spread = %v, want 1", got)
	}
	if got := Spread([]float64{5, 100, 10}); !near(got, 20) {
		t.Errorf("Spread = %v, want 20", got)
	}
}

func TestSpreadGuards(t *testing.T) {
	if got := Spread(nil); got != 0 {
		t.Errorf("Spread(nil) = %v, want 0", got)
	}
	if got := Spread([]float64{0, 5}); got != 0 {
		t.Errorf("Spread with a zero minimum = %v, want 0", got)
	}
}
