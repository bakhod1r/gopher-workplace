package histogram

import "testing"

func TestObserveAndCount(t *testing.T) {
	h := NewHistogram([]int{10, 100, 1000})
	h.Observe(5)
	h.Observe(50)
	h.Observe(500)

	if h.Count() != 3 {
		t.Errorf("Count = %d, want 3", h.Count())
	}
	if h.counts[0] != 1 || h.counts[1] != 1 || h.counts[2] != 1 {
		t.Errorf("counts = %v, want one per bucket", h.counts)
	}
}

func TestBoundIsInclusive(t *testing.T) {
	h := NewHistogram([]int{10, 100})
	h.Observe(10)
	if h.counts[0] != 1 {
		t.Errorf("a value equal to the bound belongs in that bucket: %v", h.counts)
	}
}

func TestOverflow(t *testing.T) {
	h := NewHistogram([]int{10})
	h.Observe(999)
	if h.overflow != 1 {
		t.Errorf("overflow = %d, want 1", h.overflow)
	}
	if h.Count() != 1 {
		t.Errorf("Count = %d, want 1", h.Count())
	}
	if got := h.Quantile(1); got != 10 {
		t.Errorf("Quantile(1) = %d, want 10 (last bound)", got)
	}
}

func TestQuantile(t *testing.T) {
	h := NewHistogram([]int{10, 100, 1000})
	h.Observe(5)
	h.Observe(50)
	h.Observe(500)

	cases := []struct {
		q    float64
		want int
	}{
		{0, 10},
		{0.33, 10},
		{0.5, 100},
		{0.9, 1000},
		{1, 1000},
	}

	for _, tc := range cases {
		if got := h.Quantile(tc.q); got != tc.want {
			t.Errorf("Quantile(%v) = %d, want %d", tc.q, got, tc.want)
		}
	}
}

func TestQuantileEmpty(t *testing.T) {
	h := NewHistogram([]int{10, 100})
	if got := h.Quantile(0.5); got != 0 {
		t.Errorf("Quantile on an empty histogram = %d, want 0", got)
	}
}

func TestConstantMemoryOverManySamples(t *testing.T) {
	h := NewHistogram([]int{10, 100, 1000, 10000})
	var r Recorder = h

	const n = 1_000_000
	for i := 0; i < n; i++ {
		r.Observe(i % 20000)
	}

	if h.Count() != n {
		t.Errorf("Count = %d, want %d", h.Count(), n)
	}
	if len(h.counts) != 4 {
		t.Errorf("bucket count = %d, want 4 regardless of sample count", len(h.counts))
	}
}

func BenchmarkObserve(b *testing.B) {
	h := NewHistogram([]int{10, 100, 1000, 10000})
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		h.Observe(i % 20000)
	}
}
