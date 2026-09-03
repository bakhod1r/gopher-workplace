package structclone

import (
	"reflect"
	"testing"
)

type pt struct {
	X, Y int
	Tags []string
}

func TestCloneCopiesFields(t *testing.T) {
	in := pt{X: 1, Y: 2}
	got := Clone(in)
	out, ok := got.(pt)
	if !ok {
		t.Fatalf("Clone returned %T, want pt", got)
	}
	if out.X != 1 || out.Y != 2 || out.Tags != nil {
		t.Errorf("Clone = %+v, want {1 2 []}", out)
	}
}

func TestCloneIsIndependentForValueFields(t *testing.T) {
	in := pt{X: 1}
	out := Clone(in).(pt)
	out.X = 99
	if in.X != 1 {
		t.Errorf("in.X = %d, want 1", in.X)
	}
}

func TestCloneIsShallow(t *testing.T) {
	in := pt{Tags: []string{"a"}}
	out := Clone(in).(pt)
	out.Tags[0] = "changed"
	if in.Tags[0] != "changed" {
		t.Error("the clone copied the slice; a shallow copy shares it")
	}
}

func TestCloneRejectsNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}, &pt{}} {
		if got := Clone(in); got != nil {
			t.Errorf("Clone(%#v) = %v, want nil", in, got)
		}
	}
}

func TestClonePreservesTheType(t *testing.T) {
	type other struct{ A string }
	got := Clone(other{A: "x"})
	if reflect.TypeOf(got) != reflect.TypeOf(other{}) {
		t.Errorf("type = %T, want other", got)
	}
}
