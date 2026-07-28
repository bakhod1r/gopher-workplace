package arraysharing

import "testing"

func TestViews(t *testing.T) {
	var arr [3]int
	a, b := Views(&arr)
	a[0] = 42
	if b[0] != 42 {
		t.Errorf("b[0]=%d want 42 (views must alias the same array)", b[0])
	}
}
