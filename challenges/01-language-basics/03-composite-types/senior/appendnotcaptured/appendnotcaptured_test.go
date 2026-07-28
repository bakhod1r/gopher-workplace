package appendnotcaptured

import (
	"reflect"
	"testing"
)

func TestDoubled(t *testing.T) {
	got := Doubled([]int{1, 2, 3})
	want := []int{2, 4, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Doubled=%v; want %v", got, want)
	}
}
