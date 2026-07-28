package setintersect

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	got := Intersect([]int{1, 2, 3, 4}, []int{2, 4, 6, 2})
	want := []int{2, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Intersect=%v; want %v", got, want)
	}
	if len(Intersect([]int{1}, []int{2})) != 0 {
		t.Error("disjoint -> empty")
	}
}
