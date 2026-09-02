package tildeconvbug

import (
	"testing"
	"time"
)

type Millis int64

func TestTotalSmall(t *testing.T) {
	if got := Total([]Millis{1, 2, 3}); got != 6 {
		t.Errorf("Total = %d, want 6", got)
	}
}

func TestTotalWideValues(t *testing.T) {
	got := Total([]Millis{3000000000, 3000000000})
	if got != 6000000000 {
		t.Errorf("Total = %d, want 6000000000", got)
	}
}

func TestTotalEmpty(t *testing.T) {
	if got := Total([]Millis{}); got != 0 {
		t.Errorf("Total = %d, want 0", got)
	}
}

func TestTotalScale(t *testing.T) {
	const n = 3000000
	vals := make([]Millis, n)
	for i := range vals {
		vals[i] = 1000
	}
	start := time.Now()
	got := Total(vals)
	if el := time.Since(start); el > 5*time.Second {
		t.Fatalf("Total over %d elements took %v, want under 5s", n, el)
	}
	if want := Millis(n) * 1000; got != want {
		t.Errorf("Total = %d, want %d", got, want)
	}
}
