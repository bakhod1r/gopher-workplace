package histmerge

import (
	"reflect"
	"testing"
)

func TestValid(t *testing.T) {
	cases := []struct {
		h    Hist
		want bool
	}{
		{Hist{[]float64{1, 2}, []int64{0, 0, 0}}, true},
		{Hist{nil, []int64{5}}, true},
		{Hist{[]float64{1, 2}, []int64{0, 0}}, false},
		{Hist{[]float64{2, 1}, []int64{0, 0, 0}}, false},
		{Hist{[]float64{1, 1}, []int64{0, 0, 0}}, false},
		{Hist{}, false},
	}
	for _, c := range cases {
		if got := Valid(c.h); got != c.want {
			t.Errorf("Valid(%v) = %v, want %v", c.h, got, c.want)
		}
	}
}

func TestMerge(t *testing.T) {
	a := Hist{[]float64{1, 5}, []int64{1, 2, 3}}
	b := Hist{[]float64{1, 5}, []int64{10, 20, 30}}
	got, ok := Merge(a, b)
	if !ok {
		t.Fatal("Merge reported failure for matching bounds")
	}
	want := Hist{[]float64{1, 5}, []int64{11, 22, 33}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Merge = %v, want %v", got, want)
	}
}

func TestMergeRejectsDifferentBounds(t *testing.T) {
	a := Hist{[]float64{1, 5}, []int64{1, 1, 1}}
	b := Hist{[]float64{1, 6}, []int64{1, 1, 1}}
	if _, ok := Merge(a, b); ok {
		t.Error("Merge accepted mismatched bounds")
	}
	c := Hist{[]float64{1}, []int64{1, 1}}
	if _, ok := Merge(a, c); ok {
		t.Error("Merge accepted a different number of bounds")
	}
}

func TestMergeRejectsInvalidInput(t *testing.T) {
	good := Hist{[]float64{1}, []int64{1, 1}}
	bad := Hist{[]float64{1}, []int64{1}}
	if _, ok := Merge(good, bad); ok {
		t.Error("Merge accepted a malformed histogram")
	}
}

func TestMergeDoesNotModifyInputs(t *testing.T) {
	a := Hist{[]float64{1, 5}, []int64{1, 2, 3}}
	b := Hist{[]float64{1, 5}, []int64{10, 20, 30}}
	got, _ := Merge(a, b)
	got.Counts[0] = 999
	got.Bounds[0] = 999
	if a.Counts[0] != 1 || b.Counts[0] != 10 || a.Bounds[0] != 1 {
		t.Errorf("an input was aliased: a = %v, b = %v", a, b)
	}
}

func TestMergePreservesTheTotal(t *testing.T) {
	a := Hist{[]float64{1, 5}, []int64{1, 2, 3}}
	b := Hist{[]float64{1, 5}, []int64{10, 20, 30}}
	got, _ := Merge(a, b)
	var total int64
	for _, c := range got.Counts {
		total += c
	}
	if total != 66 {
		t.Errorf("merged total = %d, want 66", total)
	}
}
