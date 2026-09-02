package scanbug

import (
	"reflect"
	"testing"
)

func TestScanValues(t *testing.T) {
	add := func(a, n int) int { return a + n }
	if got := Scan([]int{1, 2, 3}, 0, add); !reflect.DeepEqual(got, []int{1, 3, 6}) {
		t.Errorf("Scan = %v, want [1 3 6]", got)
	}
	mul := func(a, n int) int { return a * n }
	if got := Scan([]int{2, 3}, 1, mul); !reflect.DeepEqual(got, []int{2, 6}) {
		t.Errorf("Scan = %v, want [2 6]", got)
	}
}

func TestScanLastMatchesFold(t *testing.T) {
	add := func(a, n int) int { return a + n }
	s := []int{4, 5, 6}
	got := Scan(s, 10, add)
	if len(got) != len(s) {
		t.Fatalf("len = %d, want %d", len(got), len(s))
	}
	want := 10
	for _, v := range s {
		want = add(want, v)
	}
	if got[len(got)-1] != want {
		t.Errorf("last element = %v, want %v (the fold result)", got[len(got)-1], want)
	}
}

func TestScanEmpty(t *testing.T) {
	add := func(a, n int) int { return a + n }
	if got := Scan([]int{}, 5, add); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Scan(empty) = %v, want []", got)
	}
}
