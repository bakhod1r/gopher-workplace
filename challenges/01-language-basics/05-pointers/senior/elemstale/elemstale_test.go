package elemstale

import "testing"

func TestFirstOf(t *testing.T) {
	got := FirstOf([]int{1, 2})
	if got[0] != 42 {
		t.Errorf("got[0]=%d want 42 (stale pointer after realloc)", got[0])
	}
}
