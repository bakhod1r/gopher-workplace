package indexedlistgen

import "testing"

func TestIndexedAppend(t *testing.T) {
	l := NewIndexed[string]()
	if !l.Append("a") {
		t.Error(`Append("a") = false, want true`)
	}
	if l.Append("a") {
		t.Error(`Append("a") twice = true, want false`)
	}
	if !l.Append("b") {
		t.Error(`Append("b") = false, want true`)
	}
	if !l.Has("a") || !l.Has("b") {
		t.Error("Has reported a missing element")
	}
	if l.Has("c") {
		t.Error(`Has("c") = true, want false`)
	}
}

func TestIndexedAt(t *testing.T) {
	l := NewIndexed[int]()
	l.Append(10)
	l.Append(20)
	if v, ok := l.At(0); v != 10 || !ok {
		t.Errorf("At(0) = %v, %v, want 10, true", v, ok)
	}
	if v, ok := l.At(1); v != 20 || !ok {
		t.Errorf("At(1) = %v, %v, want 20, true", v, ok)
	}
	if v, ok := l.At(5); v != 0 || ok {
		t.Errorf("At(5) = %v, %v, want 0, false", v, ok)
	}
	if _, ok := l.At(-1); ok {
		t.Error("At(-1) reported ok")
	}
}
