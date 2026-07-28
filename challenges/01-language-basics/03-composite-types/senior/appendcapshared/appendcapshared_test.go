package appendcapshared

import (
	"reflect"
	"testing"
)

func TestFirstTwoPlus(t *testing.T) {
	xs := make([]int, 3, 5) // len 3, spare capacity
	xs[0], xs[1], xs[2] = 1, 2, 3
	got := FirstTwoPlus(xs, 99)
	if !reflect.DeepEqual(got, []int{1, 2, 99}) {
		t.Errorf("got %v; want [1 2 99]", got)
	}
	if xs[2] != 3 {
		t.Errorf("source clobbered: xs[2]=%d; want 3", xs[2])
	}
}
