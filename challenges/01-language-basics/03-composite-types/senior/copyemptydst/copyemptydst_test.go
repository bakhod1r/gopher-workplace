package copyemptydst

import (
	"reflect"
	"testing"
)

func TestClone(t *testing.T) {
	xs := []int{1, 2, 3}
	got := Clone(xs)
	if !reflect.DeepEqual(got, xs) {
		t.Errorf("Clone=%v; want %v", got, xs)
	}
	got[0] = 9
	if xs[0] != 1 {
		t.Error("not independent")
	}
}
