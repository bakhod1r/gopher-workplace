package firstlast

import "testing"

func TestFirstLast(t *testing.T) {
	if _, _, ok := FirstLast(nil); ok {
		t.Errorf("empty slice: ok should be false")
	}
	f, l, ok := FirstLast([]int{7})
	if f != 7 || l != 7 || !ok {
		t.Errorf("[7]=>%d,%d,%v want 7,7,true", f, l, ok)
	}
	f, l, ok = FirstLast([]int{2, 4, 6, 8})
	if f != 2 || l != 8 || !ok {
		t.Errorf("=>%d,%d,%v want 2,8,true", f, l, ok)
	}
}
