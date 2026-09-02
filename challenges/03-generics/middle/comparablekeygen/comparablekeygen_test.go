package comparablekeygen

import "testing"

func TestIndex(t *testing.T) {
	got := Index([]string{"a", "b"}, []int{1, 2})
	if len(got) != 2 || got["a"] != 1 || got["b"] != 2 {
		t.Errorf("Index = %v, want {a:1 b:2}", got)
	}
}

func TestIndexStopsAtShorter(t *testing.T) {
	got := Index([]string{"a", "b", "c"}, []int{1})
	if len(got) != 1 || got["a"] != 1 {
		t.Errorf("Index = %v, want {a:1}", got)
	}
	empty := Index([]string{"a"}, []int{})
	if empty == nil || len(empty) != 0 {
		t.Errorf("Index = %v, want an empty non-nil map", empty)
	}
}

func TestIndexAnyPanicsOnUncomparableKey(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("IndexAny with a slice key did not panic; the comparison with Index is the point")
		}
	}()
	IndexAny([]any{[]int{1}}, []any{1})
}
