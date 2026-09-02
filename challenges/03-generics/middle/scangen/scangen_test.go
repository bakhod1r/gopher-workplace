package scangen

import (
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {
	add := func(a, n int) int { return a + n }
	if got := Scan([]int{1, 2, 3}, 0, add); !reflect.DeepEqual(got, []int{1, 3, 6}) {
		t.Errorf("Scan = %v, want [1 3 6]", got)
	}
	if got := Scan([]int{}, 5, add); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Scan(empty) = %v, want [] (init is not emitted)", got)
	}
	mul := func(a, n int) int { return a * n }
	if got := Scan([]int{2}, 1, mul); !reflect.DeepEqual(got, []int{2}) {
		t.Errorf("Scan = %v, want [2]", got)
	}
}

func TestScanChangesType(t *testing.T) {
	concat := func(a string, n int) string { return a + string(rune('0'+n)) }
	got := Scan([]int{1, 2}, "", concat)
	if !reflect.DeepEqual(got, []string{"1", "12"}) {
		t.Errorf("Scan = %v, want [1 12]", got)
	}
}
