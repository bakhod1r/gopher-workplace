package thresholdbug

import (
	"reflect"
	"testing"
)

func TestAboveThreshold(t *testing.T) {
	got := AboveThreshold([]int{1, 5, 5, 8, 3}, 5)
	if !reflect.DeepEqual(got, []int{8}) {
		t.Errorf("=%v want [8] (5 is not strictly greater)", got)
	}
}
