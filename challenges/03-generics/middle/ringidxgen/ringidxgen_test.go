package ringidxgen

import (
	"reflect"
	"testing"
)

func TestRingOverwrites(t *testing.T) {
	r := NewRing[int](2)
	if got := r.Items(); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Items() = %v, want []", got)
	}
	r.Add(1)
	r.Add(2)
	if got := r.Items(); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Items() = %v, want [1 2]", got)
	}
	r.Add(3)
	if got := r.Items(); !reflect.DeepEqual(got, []int{2, 3}) {
		t.Errorf("Items() = %v, want [2 3]", got)
	}
	r.Add(4)
	r.Add(5)
	if got := r.Items(); !reflect.DeepEqual(got, []int{4, 5}) {
		t.Errorf("Items() = %v, want [4 5]", got)
	}
}

func TestRingPartial(t *testing.T) {
	r := NewRing[string](3)
	r.Add("a")
	if got := r.Items(); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("Items() = %v, want [a]", got)
	}
}

func TestRingZeroSize(t *testing.T) {
	r := NewRing[int](0)
	r.Add(1)
	if got := r.Items(); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Items() = %v, want []", got)
	}
}
