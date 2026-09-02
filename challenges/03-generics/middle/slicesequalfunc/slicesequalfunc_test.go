package slicesequalfunc

import (
	"strconv"
	"testing"
)

func TestSameRows(t *testing.T) {
	matches := func(n int, s string) bool { return strconv.Itoa(n) == s }
	if !SameRows([]int{1, 2}, []string{"1", "2"}, matches) {
		t.Error("SameRows = false, want true")
	}
	if SameRows([]int{1}, []string{"2"}, matches) {
		t.Error("SameRows = true, want false")
	}
	if SameRows([]int{1}, []string{"1", "2"}, matches) {
		t.Error("SameRows with different lengths = true, want false")
	}
	if !SameRows([]int(nil), []string{}, matches) {
		t.Error("SameRows(nil, []) = false, want true")
	}
}
