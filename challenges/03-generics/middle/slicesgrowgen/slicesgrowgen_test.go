package slicesgrowgen

import (
	"reflect"
	"testing"
)

func TestCollect(t *testing.T) {
	if got := Collect([]int{1}, 2, 3); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Collect = %v, want [1 2 3]", got)
	}
	if got := Collect([]int(nil), 1); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Collect(nil, 1) = %v, want [1]", got)
	}
	if got := Collect([]int{1}); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Collect with no values = %v, want [1]", got)
	}
}

func TestCollectReservesCapacity(t *testing.T) {
	got := Collect(make([]int, 0), 1, 2, 3, 4, 5)
	if cap(got) < 5 {
		t.Errorf("cap = %d, want at least 5", cap(got))
	}
	if len(got) != 5 {
		t.Errorf("len = %d, want 5", len(got))
	}
}
