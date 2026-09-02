package betweengen

import "testing"

func TestBetween(t *testing.T) {
	if !Between(2, 1, 3) {
		t.Error("Between(2, 1, 3) = false, want true")
	}
	if !Between(1, 1, 3) {
		t.Error("Between(1, 1, 3) = false, want true (bounds are inclusive)")
	}
	if !Between(3, 1, 3) {
		t.Error("Between(3, 1, 3) = false, want true (bounds are inclusive)")
	}
	if Between(4, 1, 3) {
		t.Error("Between(4, 1, 3) = true, want false")
	}
	if Between(0, 1, 3) {
		t.Error("Between(0, 1, 3) = true, want false")
	}
	if !Between("b", "a", "c") {
		t.Error(`Between("b", "a", "c") = false, want true`)
	}
}
