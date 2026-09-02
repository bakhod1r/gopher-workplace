package swapgen

import "testing"

func TestSwap(t *testing.T) {
	if a, b := Swap(1, 2); a != 2 || b != 1 {
		t.Errorf("Swap(1, 2) = %v, %v, want 2, 1", a, b)
	}
	if a, b := Swap("a", "b"); a != "b" || b != "a" {
		t.Errorf("Swap(\"a\", \"b\") = %q, %q, want \"b\", \"a\"", a, b)
	}
	if a, b := Swap(true, false); a != false || b != true {
		t.Errorf("Swap(true, false) = %v, %v, want false, true", a, b)
	}
}
