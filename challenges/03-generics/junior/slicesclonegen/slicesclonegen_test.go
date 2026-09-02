package slicesclonegen

import (
	"reflect"
	"testing"
)

func TestDetach(t *testing.T) {
	in := []int{1, 2}
	got := Detach(in)
	if !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Detach = %v, want [1 2]", got)
	}
	got[0] = 99
	if in[0] != 1 {
		t.Errorf("writing to the copy changed the input: in[0] = %v, want 1", in[0])
	}
}

func TestDetachNil(t *testing.T) {
	got := Detach([]string(nil))
	if got == nil {
		t.Error("Detach(nil) = nil, want an empty non-nil slice")
	}
	if len(got) != 0 {
		t.Errorf("Detach(nil) = %v, want []", got)
	}
}
