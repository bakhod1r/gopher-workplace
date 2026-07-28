package shrinkretain

import "testing"

func TestPop(t *testing.T) {
	a, b := 1, 2
	s := []*int{&a, &b}
	Pop(&s)
	// The backing array slot [1] must be nil so the object can be GC'd.
	full := s[:2]
	if full[1] != nil {
		t.Errorf("vacated slot still holds a pointer (retention leak)")
	}
}
