package foldcollapse

import (
	"reflect"
	"testing"
)

func TestCollapse(t *testing.T) {
	got := Collapse([]string{"a", "b", "b", "b", "c"})
	if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("Collapse = %v, want [a b c]", got)
	}
}

func TestCollapseKeepsNonAdjacentRepeats(t *testing.T) {
	got := Collapse([]string{"a", "b", "a", "b"})
	if !reflect.DeepEqual(got, []string{"a", "b", "a", "b"}) {
		t.Errorf("Collapse = %v, want the stack unchanged", got)
	}
}

func TestCollapseMutualRecursion(t *testing.T) {
	got := Collapse([]string{"a", "a", "b", "b", "a", "a"})
	if !reflect.DeepEqual(got, []string{"a", "b", "a"}) {
		t.Errorf("Collapse = %v, want [a b a]", got)
	}
}

func TestCollapseEdgeCases(t *testing.T) {
	if got := Collapse([]string{"a"}); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("Collapse = %v, want [a]", got)
	}
	got := Collapse(nil)
	if got == nil || len(got) != 0 {
		t.Errorf("Collapse(nil) = %v, want empty non-nil slice", got)
	}
}

func TestCollapseDoesNotModifyInput(t *testing.T) {
	in := []string{"a", "b", "b", "c"}
	before := append([]string(nil), in...)
	Collapse(in)
	if !reflect.DeepEqual(in, before) {
		t.Errorf("input changed: %v, want %v", in, before)
	}
}

func TestDepth(t *testing.T) {
	cases := []struct {
		stack []string
		want  int
	}{
		{[]string{"a", "b", "b", "b"}, 2},
		{[]string{"a", "b", "a", "b"}, 0},
		{nil, 0},
		{[]string{"r", "r", "r", "r", "r"}, 4},
	}
	for _, c := range cases {
		if got := Depth(c.stack); got != c.want {
			t.Errorf("Depth(%v) = %d, want %d", c.stack, got, c.want)
		}
	}
}
