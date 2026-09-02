package modegen

import "testing"

func TestMode(t *testing.T) {
	if v, ok := Mode([]int{1, 2, 2}); v != 2 || !ok {
		t.Errorf("Mode = %v, %v, want 2, true", v, ok)
	}
	if v, ok := Mode([]int{5}); v != 5 || !ok {
		t.Errorf("Mode = %v, %v, want 5, true", v, ok)
	}
	if v, ok := Mode([]int{}); v != 0 || ok {
		t.Errorf("Mode(empty) = %v, %v, want 0, false", v, ok)
	}
	if v, ok := Mode([]string{"a", "b", "b"}); v != "b" || !ok {
		t.Errorf("Mode = %q, %v, want b, true", v, ok)
	}
}

func TestModeTieIsDeterministic(t *testing.T) {
	for i := 0; i < 50; i++ {
		if v, _ := Mode([]int{1, 1, 2, 2}); v != 1 {
			t.Fatalf("Mode = %v, want 1 (earliest wins ties, every run)", v)
		}
	}
}
