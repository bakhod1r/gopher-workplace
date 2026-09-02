package mycmpgen

import "testing"

type Celsius float64

func TestLargest(t *testing.T) {
	if v, ok := Largest([]int{1, 9, 3}); v != 9 || !ok {
		t.Errorf("Largest = %v, %v, want 9, true", v, ok)
	}
	if v, ok := Largest([]string{"a", "c", "b"}); v != "c" || !ok {
		t.Errorf("Largest = %q, %v, want c, true", v, ok)
	}
	if v, ok := Largest([]uint8{1, 2}); v != 2 || !ok {
		t.Errorf("Largest = %v, %v, want 2, true", v, ok)
	}
}

func TestLargestNamedType(t *testing.T) {
	if v, ok := Largest([]Celsius{1, 5}); v != Celsius(5) || !ok {
		t.Errorf("Largest = %v, %v, want 5, true (the ~ admits named types)", v, ok)
	}
}

func TestLargestEmpty(t *testing.T) {
	if v, ok := Largest([]float64{}); v != 0 || ok {
		t.Errorf("Largest(empty) = %v, %v, want 0, false", v, ok)
	}
}
