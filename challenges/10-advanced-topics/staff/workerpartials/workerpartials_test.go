package workerpartials

import (
	"reflect"
	"testing"
)

func TestHistogramSmall(t *testing.T) {
	if got := Histogram([]int{0, 1, 2, 3}, 2, 2); !reflect.DeepEqual(got, []int64{2, 2}) {
		t.Errorf("Histogram = %v, want [2 2]", got)
	}
	if got := Histogram(nil, 3, 4); !reflect.DeepEqual(got, []int64{0, 0, 0}) {
		t.Errorf("Histogram = %v, want [0 0 0]", got)
	}
	if got := Histogram([]int{1}, 0, 2); got != nil {
		t.Errorf("Histogram = %v, want nil", got)
	}
}

func TestHistogramNegativeValues(t *testing.T) {
	got := Histogram([]int{-1, -2, -3}, 3, 2)
	if !reflect.DeepEqual(got, []int64{1, 1, 1}) {
		t.Errorf("Histogram = %v, want [1 1 1]: negative values must land in range", got)
	}
}

func TestHistogramMatchesSerial(t *testing.T) {
	const buckets = 13
	data := make([]int, 100003)
	want := make([]int64, buckets)
	for i := range data {
		data[i] = i * 7
		want[data[i]%buckets]++
	}
	for _, w := range []int{1, 2, 3, 8, 32, 1 << 20} {
		got := Histogram(data, buckets, w)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("workers=%d: Histogram = %v, want %v", w, got, want)
		}
	}
}

func TestHistogramRepeatable(t *testing.T) {
	data := make([]int, 50000)
	for i := range data {
		data[i] = i
	}
	first := Histogram(data, 7, 8)
	for i := 0; i < 20; i++ {
		if got := Histogram(data, 7, 8); !reflect.DeepEqual(got, first) {
			t.Fatalf("round %d: %v, want %v: the workers are sharing counters", i, got, first)
		}
	}
}
