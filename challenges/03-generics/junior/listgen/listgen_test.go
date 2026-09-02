package listgen

import "testing"

func TestList(t *testing.T) {
	var l List[int]
	if l.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", l.Len())
	}
	l.Append(1)
	l.Append(2)
	if v, ok := l.At(0); v != 1 || !ok {
		t.Errorf("At(0) = %v, %v, want 1, true", v, ok)
	}
	if v, ok := l.At(1); v != 2 || !ok {
		t.Errorf("At(1) = %v, %v, want 2, true", v, ok)
	}
	if v, ok := l.At(5); v != 0 || ok {
		t.Errorf("At(5) = %v, %v, want 0, false", v, ok)
	}
	if v, ok := l.At(-1); v != 0 || ok {
		t.Errorf("At(-1) = %v, %v, want 0, false", v, ok)
	}
	if l.Len() != 2 {
		t.Errorf("Len() = %d, want 2", l.Len())
	}
}
