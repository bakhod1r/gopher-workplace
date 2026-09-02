package percentilerankbug

import (
	"testing"
	"time"
)

func TestPercentileMedian(t *testing.T) {
	got, ok := Percentile([]int{5, 1, 4, 2, 3}, 50)
	if !ok || got != 3 {
		t.Errorf("Percentile = %v, %v; want 3, true", got, ok)
	}
}

func TestPercentileLowRank(t *testing.T) {
	got, ok := Percentile([]int{1, 2, 3, 4}, 25)
	if !ok || got != 1 {
		t.Errorf("Percentile = %v, %v; want 1, true", got, ok)
	}
}

func TestPercentileEnds(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	if got, ok := Percentile(xs, 0); !ok || got != 1 {
		t.Errorf("Percentile p0 = %v, %v; want 1, true", got, ok)
	}
	if got, ok := Percentile(xs, 100); !ok || got != 5 {
		t.Errorf("Percentile p100 = %v, %v; want 5, true", got, ok)
	}
}

func TestPercentileInvalid(t *testing.T) {
	if _, ok := Percentile([]int{1, 2, 3}, 101); ok {
		t.Errorf("ok = true, want false")
	}
	if _, ok := Percentile([]int{}, 50); ok {
		t.Errorf("ok = true, want false")
	}
}

func TestPercentileDoesNotMutate(t *testing.T) {
	xs := []int{3, 1, 2}
	Percentile(xs, 50)
	if xs[0] != 3 || xs[1] != 1 || xs[2] != 2 {
		t.Errorf("input mutated: %v", xs)
	}
}

func TestPercentileScale(t *testing.T) {
	const n = 2000001
	xs := make([]int, n)
	for i := range xs {
		xs[i] = (i * 7919) % n
	}
	start := time.Now()
	got, ok := Percentile(xs, 50)
	if el := time.Since(start); el > 10*time.Second {
		t.Fatalf("Percentile over %d elements took %v, want under 10s", n, el)
	}
	if !ok || got != (n-1)/2 {
		t.Errorf("Percentile = %v, %v; want %d, true", got, ok, (n-1)/2)
	}
}
