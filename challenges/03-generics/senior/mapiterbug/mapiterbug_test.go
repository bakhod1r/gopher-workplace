package mapiterbug

import "testing"

func TestModeTieIsDeterministic(t *testing.T) {
	for i := 0; i < 100; i++ {
		v, ok := Mode([]int{1, 1, 2, 2})
		if !ok || v != 1 {
			t.Fatalf("Mode = %v, %v, want 1, true on every run (earliest wins ties)", v, ok)
		}
	}
}

func TestModeClearWinner(t *testing.T) {
	if v, ok := Mode([]int{1, 2, 2}); v != 2 || !ok {
		t.Errorf("Mode = %v, %v, want 2, true", v, ok)
	}
	if v, ok := Mode([]string{"a", "b", "b"}); v != "b" || !ok {
		t.Errorf("Mode = %q, %v, want b, true", v, ok)
	}
}

func TestModeEmpty(t *testing.T) {
	if v, ok := Mode([]int{}); v != 0 || ok {
		t.Errorf("Mode(empty) = %v, %v, want 0, false", v, ok)
	}
}
