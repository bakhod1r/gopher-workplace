package insertat

import (
	"reflect"
	"testing"
)

func TestInsertAt(t *testing.T) {
	if got := InsertAt([]int{1, 3}, 1, 2); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt = %v, want [1 2 3]", got)
	}
	if got := InsertAt([]int{2, 3}, 0, 1); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt = %v, want [1 2 3]", got)
	}
	if got := InsertAt([]int{1, 2}, 2, 3); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt = %v, want [1 2 3]", got)
	}
}

func TestInsertAtClamps(t *testing.T) {
	if got := InsertAt([]int{1, 2}, 99, 3); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt = %v, want [1 2 3]", got)
	}
	if got := InsertAt([]int{2, 3}, -5, 1); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt = %v, want [1 2 3]", got)
	}
}

func TestInsertAtEmpty(t *testing.T) {
	if got := InsertAt(nil, 0, 7); !reflect.DeepEqual(got, []int{7}) {
		t.Errorf("InsertAt = %v, want [7]", got)
	}
}

func TestInsertAtLongShift(t *testing.T) {
	s := make([]int, 0, 100)
	for i := 0; i < 50; i++ {
		s = append(s, i)
	}
	got := InsertAt(s, 0, -1)
	if got[0] != -1 {
		t.Fatalf("got[0] = %d, want -1", got[0])
	}
	for i := 0; i < 50; i++ {
		if got[i+1] != i {
			t.Fatalf("got[%d] = %d, want %d: the shift lost an element", i+1, got[i+1], i)
		}
	}
}

func TestInsertAtReusesCapacity(t *testing.T) {
	s := make([]int, 2, 8)
	if n := testing.AllocsPerRun(100, func() { _ = InsertAt(s[:2], 1, 9) }); n != 0 {
		t.Errorf("InsertAt made %v allocations, want 0 when the capacity allows", n)
	}
}
