package unsigneddeltabug

import (
	"reflect"
	"testing"
	"time"
)

func TestDeltasReset(t *testing.T) {
	got := Deltas([]uint8{10, 250, 5})
	if !reflect.DeepEqual(got, []uint8{10, 240, 5}) {
		t.Errorf("Deltas = %v, want [10 240 5]", got)
	}
}

func TestDeltasMonotonic(t *testing.T) {
	got := Deltas([]uint64{5, 9, 20})
	if !reflect.DeepEqual(got, []uint64{5, 4, 11}) {
		t.Errorf("Deltas = %v, want [5 4 11]", got)
	}
}

func TestDeltasEmpty(t *testing.T) {
	if got := Deltas([]uint32{}); len(got) != 0 {
		t.Errorf("Deltas = %v, want []", got)
	}
}

func TestDeltasScale(t *testing.T) {
	const n = 2000000
	samples := make([]uint32, n)
	for i := range samples {
		samples[i] = uint32(i % 1000)
	}
	start := time.Now()
	got := Deltas(samples)
	if el := time.Since(start); el > 5*time.Second {
		t.Fatalf("Deltas over %d elements took %v, want under 5s", n, el)
	}
	if len(got) != n {
		t.Fatalf("Deltas len = %d, want %d", len(got), n)
	}
	var sum uint64
	for _, d := range got {
		sum += uint64(d)
	}
	if want := uint64(n - n/1000); sum != want {
		t.Errorf("sum of deltas = %d, want %d", sum, want)
	}
}
