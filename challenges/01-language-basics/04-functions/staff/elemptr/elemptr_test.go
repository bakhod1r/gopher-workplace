package elemptr

import "testing"

func TestBumpHead(t *testing.T) {
	xs := BumpHead(9)
	if xs[0] != 42 {
		t.Errorf("xs[0]=%d want 42 (stale pointer after realloc)", xs[0])
	}
}
