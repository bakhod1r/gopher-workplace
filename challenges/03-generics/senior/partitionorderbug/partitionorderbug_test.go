package partitionorderbug

import (
	"reflect"
	"testing"
)

func even(n int) bool { return n%2 == 0 }

func TestPartitionKeepsOrder(t *testing.T) {
	yes, no := Partition([]int{1, 2, 3, 4}, even)
	if !reflect.DeepEqual(yes, []int{2, 4}) {
		t.Errorf("yes = %v, want [2 4]", yes)
	}
	if !reflect.DeepEqual(no, []int{1, 3}) {
		t.Errorf("no = %v, want [1 3]", no)
	}
}

func TestPartitionAllOneSide(t *testing.T) {
	yes, no := Partition([]int{1, 3, 5}, even)
	if len(yes) != 0 {
		t.Errorf("yes = %v, want []", yes)
	}
	if !reflect.DeepEqual(no, []int{1, 3, 5}) {
		t.Errorf("no = %v, want [1 3 5]", no)
	}
}

func TestPartitionEmpty(t *testing.T) {
	yes, no := Partition([]int{}, even)
	if len(yes) != 0 || len(no) != 0 {
		t.Errorf("Partition = %v, %v, want empty", yes, no)
	}
}
