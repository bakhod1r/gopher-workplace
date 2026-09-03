package clipbug

import (
	"reflect"
	"testing"
)

func TestHead(t *testing.T) {
	if got := Head([]int{1, 2, 3}, 2); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Head = %v, want [1 2]", got)
	}
	if got := Head([]int{1}, 9); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Head = %v, want [1]", got)
	}
	if got := Head([]int{1}, -1); len(got) != 0 {
		t.Errorf("Head = %v, want []", got)
	}
}

func TestHeadIsSafeToAppendTo(t *testing.T) {
	s := []int{1, 2, 3, 4}
	got := append(Head(s, 2), 99)
	if s[2] != 3 {
		t.Errorf("appending to the result overwrote the source: s = %v, want [1 2 3 4]", s)
	}
	if !reflect.DeepEqual(got, []int{1, 2, 99}) {
		t.Errorf("append result = %v, want [1 2 99]", got)
	}
}
