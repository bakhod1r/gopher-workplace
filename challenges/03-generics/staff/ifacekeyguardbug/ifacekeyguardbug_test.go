package ifacekeyguardbug

import (
	"reflect"
	"testing"
)

func TestDistinctComparable(t *testing.T) {
	got := Distinct([]any{1, 1, 2})
	if !reflect.DeepEqual(got, []any{1, 2}) {
		t.Errorf("Distinct = %v, want [1 2]", got)
	}
}

func TestDistinctUncomparableDoesNotPanic(t *testing.T) {
	got := Distinct([]any{1, []int{2}, 1})
	if !reflect.DeepEqual(got, []any{1, []int{2}}) {
		t.Errorf("Distinct = %v, want [1 [2]]", got)
	}
}

func TestDistinctOnlyUncomparable(t *testing.T) {
	got := Distinct([]any{map[string]int{"a": 1}, []string{"b"}})
	if len(got) != 2 {
		t.Errorf("Distinct len = %d, want 2", len(got))
	}
}

func TestDistinctConcreteType(t *testing.T) {
	got := Distinct([]string{"a", "b", "a"})
	if !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("Distinct = %v, want [a b]", got)
	}
}
