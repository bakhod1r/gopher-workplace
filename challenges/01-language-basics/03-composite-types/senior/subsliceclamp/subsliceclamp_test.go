package subsliceclamp

import (
	"reflect"
	"testing"
)

func TestTake(t *testing.T) {
	xs := []int{1, 2, 3}
	if got := Take(xs, 2); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Take 2=%v", got)
	}
	if got := Take(xs, 10); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Take 10 must clamp, got %v", got)
	}
	if got := Take(xs, 0); len(got) != 0 {
		t.Errorf("Take 0=%v", got)
	}
}
