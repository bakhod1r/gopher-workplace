package pushnoassign

import "testing"

func TestPush(t *testing.T) {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if len(s.data) != 3 {
		t.Errorf("len=%d want 3", len(s.data))
	}
}
