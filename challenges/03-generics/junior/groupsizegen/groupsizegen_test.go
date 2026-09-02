package groupsizegen

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	yes, no := Partition([]int{1, 2, 3}, isEven)
	if !reflect.DeepEqual(yes, []int{2}) {
		t.Errorf("accepted = %v, want [2]", yes)
	}
	if !reflect.DeepEqual(no, []int{1, 3}) {
		t.Errorf("rejected = %v, want [1 3]", no)
	}

	nonEmpty := func(s string) bool { return s != "" }
	sy, sn := Partition([]string{"a"}, nonEmpty)
	if !reflect.DeepEqual(sy, []string{"a"}) || !reflect.DeepEqual(sn, []string{}) {
		t.Errorf("Partition([a]) = %v, %v, want [a], []", sy, sn)
	}

	ey, en := Partition([]int{}, isEven)
	if !reflect.DeepEqual(ey, []int{}) || !reflect.DeepEqual(en, []int{}) {
		t.Errorf("Partition([]) = %v, %v, want [], []", ey, en)
	}
}
