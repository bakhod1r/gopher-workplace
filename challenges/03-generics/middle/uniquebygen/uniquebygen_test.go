package uniquebygen

import (
	"reflect"
	"testing"
)

func TestUniqueBy(t *testing.T) {
	mod10 := func(n int) int { return n % 10 }
	if got := UniqueBy([]int{1, 11, 2}, mod10); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("UniqueBy = %v, want [1 2] (first wins)", got)
	}
	if got := UniqueBy([]int{}, mod10); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("UniqueBy(empty) = %v, want []", got)
	}
}

func TestUniqueByStructs(t *testing.T) {
	type user struct {
		id   int
		tags []string
	}
	users := []user{{1, nil}, {1, []string{"x"}}, {2, nil}}
	got := UniqueBy(users, func(u user) int { return u.id })
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	if got[0].tags != nil {
		t.Errorf("kept the wrong element for id 1: %+v", got[0])
	}
}
