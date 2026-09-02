package groupbygen

import (
	"reflect"
	"testing"
)

func TestGroupBy(t *testing.T) {
	parity := func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	}
	got := GroupBy([]int{1, 2, 3, 4}, parity)
	if !reflect.DeepEqual(got["odd"], []int{1, 3}) {
		t.Errorf(`got["odd"] = %v, want [1 3]`, got["odd"])
	}
	if !reflect.DeepEqual(got["even"], []int{2, 4}) {
		t.Errorf(`got["even"] = %v, want [2 4]`, got["even"])
	}
	if len(got) != 2 {
		t.Errorf("len = %d, want 2", len(got))
	}
}

func TestGroupByStructs(t *testing.T) {
	type row struct {
		name string
		n    int
	}
	got := GroupBy([]row{{"a", 1}, {"b", 1}, {"c", 2}}, func(r row) int { return r.n })
	if len(got[1]) != 2 || len(got[2]) != 1 {
		t.Errorf("grouping by n is wrong: %v", got)
	}
}

func TestGroupByEmpty(t *testing.T) {
	got := GroupBy([]int{}, func(n int) int { return n })
	if got == nil || len(got) != 0 {
		t.Errorf("GroupBy(empty) = %v, want an empty non-nil map", got)
	}
}
