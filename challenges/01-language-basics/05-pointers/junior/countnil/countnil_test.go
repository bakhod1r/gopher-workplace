package countnil

import "testing"

func TestCountNil(t *testing.T) {
	a := 1
	if got := CountNil([]*int{&a, nil, nil, &a}); got != 2 {
		t.Errorf("=%d want 2", got)
	}
}
