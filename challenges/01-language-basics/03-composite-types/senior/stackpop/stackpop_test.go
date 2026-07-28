package stackpop

import (
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	s := []int{1, 2, 3}
	s, v, ok := Pop(s)
	if !ok || v != 3 || !reflect.DeepEqual(s, []int{1, 2}) {
		t.Errorf("pop1: s=%v v=%d ok=%v", s, v, ok)
	}
	s, v, _ = Pop(s)
	if v != 2 || !reflect.DeepEqual(s, []int{1}) {
		t.Errorf("pop2: s=%v v=%d", s, v)
	}
	_, _, ok = Pop(nil)
	if ok {
		t.Error("pop empty should be false")
	}
}
