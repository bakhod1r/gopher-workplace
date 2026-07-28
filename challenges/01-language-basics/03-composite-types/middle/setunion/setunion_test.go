package setunion

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	got := Union([]int{3, 1, 2}, []int{2, 4, 1})
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Union=%v; want %v", got, want)
	}
	if len(Union(nil, nil)) != 0 {
		t.Error("empty union")
	}
}
