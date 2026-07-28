package slicedata

import "testing"

func TestSetFirst(t *testing.T) {
	s := []int{1, 2, 3}
	SetFirst(s, 42)
	if s[0] != 42 {
		t.Errorf("s[0]=%d want 42 (wrote to the header, not the data)", s[0])
	}
}
