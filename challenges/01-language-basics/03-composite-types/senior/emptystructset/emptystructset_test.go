package emptystructset

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	got := Intersect([]int{1, 2, 3, 2}, []int{2, 3, 5})
	want := []int{2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Intersect=%v; want %v", got, want)
	}
}
