package slicesmaxmin

import "testing"

func TestPeak(t *testing.T) {
	if v, ok := Peak([]int{1, 9, 3}); v != 9 || !ok {
		t.Errorf("Peak = %v, %v, want 9, true", v, ok)
	}
	if v, ok := Peak([]int{}); v != 0 || ok {
		t.Errorf("Peak([]int{}) = %v, %v, want 0, false (must not panic)", v, ok)
	}
}

func TestFloor(t *testing.T) {
	if v, ok := Floor([]int{4, 1, 3}); v != 1 || !ok {
		t.Errorf("Floor = %v, %v, want 1, true", v, ok)
	}
	if v, ok := Floor([]string{}); v != "" || ok {
		t.Errorf("Floor(empty) = %q, %v, want \"\", false", v, ok)
	}
	if v, ok := Floor([]string{"b", "a"}); v != "a" || !ok {
		t.Errorf("Floor = %q, %v, want a, true", v, ok)
	}
}
