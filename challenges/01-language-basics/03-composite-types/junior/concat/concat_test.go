package concat

import (
	"reflect"
	"testing"
)

func TestConcat(t *testing.T) {
	got := Concat([]int{1, 2}, []int{3}, nil, []int{4, 5})
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Concat=%v; want %v", got, want)
	}
	if len(Concat()) != 0 {
		t.Error("no args -> empty")
	}
}
