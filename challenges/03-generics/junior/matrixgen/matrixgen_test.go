package matrixgen

import "testing"

func TestMatrix(t *testing.T) {
	m := NewMatrix[int](2, 3)
	if v, ok := m.At(0, 0); v != 0 || !ok {
		t.Errorf("At(0, 0) = %v, %v, want 0, true", v, ok)
	}
	if !m.Set(1, 2, 5) {
		t.Fatal("Set(1, 2, 5) = false, want true")
	}
	if v, ok := m.At(1, 2); v != 5 || !ok {
		t.Errorf("At(1, 2) = %v, %v, want 5, true", v, ok)
	}
	if v, ok := m.At(0, 0); v != 0 || !ok {
		t.Errorf("At(0, 0) = %v, %v, want 0, true (other cells untouched)", v, ok)
	}
}

func TestMatrixBounds(t *testing.T) {
	m := NewMatrix[string](2, 2)
	if _, ok := m.At(9, 0); ok {
		t.Error("At(9, 0) reported ok, want false")
	}
	if _, ok := m.At(0, -1); ok {
		t.Error("At(0, -1) reported ok, want false")
	}
	if m.Set(-1, 0, "x") {
		t.Error("Set(-1, 0) = true, want false")
	}
	if m.Set(0, 2, "x") {
		t.Error("Set(0, 2) on a 2x2 matrix = true, want false")
	}
}
