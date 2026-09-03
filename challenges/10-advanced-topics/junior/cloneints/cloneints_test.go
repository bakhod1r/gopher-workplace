package cloneints

import (
	"reflect"
	"testing"
)

func TestCloneContents(t *testing.T) {
	if got := Clone([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Clone = %v, want [1 2 3]", got)
	}
	if got := Clone(nil); len(got) != 0 {
		t.Errorf("Clone(nil) = %v, want empty", got)
	}
}

func TestCloneIsIndependent(t *testing.T) {
	s := []int{1, 2, 3}
	c := Clone(s)
	s[0] = 99
	if c[0] != 1 {
		t.Errorf("c[0] = %d, want 1: the clone still views s", c[0])
	}
	c[1] = 42
	if s[1] != 2 {
		t.Errorf("s[1] = %d, want 2: s still views the clone", s[1])
	}
}
