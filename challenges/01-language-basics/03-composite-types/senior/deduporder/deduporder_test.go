package deduporder

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	got := Unique([]int{3, 1, 3, 2, 1, 4})
	want := []int{3, 1, 2, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unique=%v; want %v", got, want)
	}
	if len(Unique(nil)) != 0 {
		t.Error("nil -> empty")
	}
}
