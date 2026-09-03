package foldjoinbug

import (
	"reflect"
	"testing"
)

func TestFoldKeepsStackOrder(t *testing.T) {
	got := Fold([]Sample{{[]string{"main", "a"}, 3}})
	if !reflect.DeepEqual(got, []string{"main;a 3"}) {
		t.Errorf("Fold = %v, want [main;a 3] — the stack order is the call path", got)
	}
}

func TestFoldKeepsDifferentPathsApart(t *testing.T) {
	got := Fold([]Sample{
		{[]string{"a", "b"}, 1},
		{[]string{"b", "a"}, 2},
	})
	want := []string{"b;a 2", "a;b 1"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Fold = %v, want %v — a calls b and b calls a are different paths", got, want)
	}
}

func TestFoldSumsIdenticalStacks(t *testing.T) {
	got := Fold([]Sample{
		{[]string{"a", "b"}, 3},
		{[]string{"a", "b"}, 2},
	})
	if !reflect.DeepEqual(got, []string{"a;b 5"}) {
		t.Errorf("Fold = %v, want [a;b 5]", got)
	}
}

func TestFoldRecursion(t *testing.T) {
	got := Fold([]Sample{{[]string{"main", "rec", "rec"}, 4}})
	if !reflect.DeepEqual(got, []string{"main;rec;rec 4"}) {
		t.Errorf("Fold = %v, want [main;rec;rec 4]", got)
	}
}

func TestFoldDropsJunk(t *testing.T) {
	got := Fold([]Sample{{[]string{"a"}, 0}, {nil, 5}})
	if got == nil || len(got) != 0 {
		t.Errorf("Fold = %v, want empty non-nil slice", got)
	}
}
