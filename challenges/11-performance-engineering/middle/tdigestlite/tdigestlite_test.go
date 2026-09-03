package tdigestlite

import "testing"

func build(bounds []float64, values ...float64) *Sketch {
	s := New(bounds)
	for _, v := range values {
		s.Add(v)
	}
	return s
}

func TestCount(t *testing.T) {
	s := build([]float64{1, 10}, 0.5, 5, 50)
	if got := s.Count(); got != 3 {
		t.Errorf("Count = %d, want 3", got)
	}
	if got := New([]float64{1}).Count(); got != 0 {
		t.Errorf("Count = %d, want 0", got)
	}
}

func TestQuantile(t *testing.T) {
	// Ten values: five at 0.5, five at 5. Bounds 1 and 10.
	s := New([]float64{1, 10})
	for i := 0; i < 5; i++ {
		s.Add(0.5)
	}
	for i := 0; i < 5; i++ {
		s.Add(5)
	}
	if got, ok := s.Quantile(50); !ok || got != 1 {
		t.Errorf("Quantile(50) = %v, %v, want 1, true", got, ok)
	}
	if got, ok := s.Quantile(51); !ok || got != 10 {
		t.Errorf("Quantile(51) = %v, %v, want 10, true", got, ok)
	}
	if got, ok := s.Quantile(100); !ok || got != 10 {
		t.Errorf("Quantile(100) = %v, %v, want 10, true", got, ok)
	}
	if got, ok := s.Quantile(0); !ok || got != 1 {
		t.Errorf("Quantile(0) = %v, %v, want 1, true", got, ok)
	}
}

func TestQuantileInTheOverflowBucket(t *testing.T) {
	s := build([]float64{1, 10}, 0.5, 1000)
	if _, ok := s.Quantile(99); ok {
		t.Error("Quantile in the overflow bucket reported a bound; the value is unbounded above")
	}
	if got, ok := s.Quantile(50); !ok || got != 1 {
		t.Errorf("Quantile(50) = %v, %v, want 1, true", got, ok)
	}
}

func TestQuantileEmpty(t *testing.T) {
	s := New([]float64{1, 10})
	if got, ok := s.Quantile(50); ok || got != 0 {
		t.Errorf("Quantile on an empty sketch = %v, %v, want 0, false", got, ok)
	}
}

func TestSketchIsFixedSize(t *testing.T) {
	s := New([]float64{1, 10, 100})
	for i := 0; i < 100_000; i++ {
		s.Add(float64(i % 200))
	}
	if s.Count() != 100_000 {
		t.Errorf("Count = %d, want 100000", s.Count())
	}
	if got, ok := s.Quantile(50); !ok || got != 100 {
		t.Errorf("Quantile(50) = %v, %v, want 100, true", got, ok)
	}
}
