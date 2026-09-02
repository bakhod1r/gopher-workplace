package slicescompactfunc

import (
	"reflect"
	"testing"
)

func TestDedupe(t *testing.T) {
	eq := func(a, b int) bool { return a == b }
	if got := Dedupe([]int{1, 1, 2, 2, 1}, eq); !reflect.DeepEqual(got, []int{1, 2, 1}) {
		t.Errorf("Dedupe = %v, want [1 2 1] (only adjacent runs collapse)", got)
	}
	if got := Dedupe([]int{}, eq); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Dedupe(empty) = %v, want []", got)
	}
}

func TestDedupeCustomEquality(t *testing.T) {
	type line struct {
		ts  int
		msg string
	}
	sameMsg := func(a, b line) bool { return a.msg == b.msg }
	got := Dedupe([]line{{1, "x"}, {2, "x"}, {3, "y"}}, sameMsg)
	if len(got) != 2 || got[0].msg != "x" || got[1].msg != "y" {
		t.Errorf("Dedupe = %+v, want two entries x, y", got)
	}
}

func TestDedupeDoesNotMutate(t *testing.T) {
	in := []int{1, 1, 2}
	Dedupe(in, func(a, b int) bool { return a == b })
	if !reflect.DeepEqual(in, []int{1, 1, 2}) {
		t.Errorf("input mutated: %v", in)
	}
}
