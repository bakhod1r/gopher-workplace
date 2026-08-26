package vmopcode

import "testing"

func TestVM(t *testing.T) {
	v := &VM{}
	if got := v.Next(); got != 1 {
		t.Errorf("got %d", got)
	}
}
