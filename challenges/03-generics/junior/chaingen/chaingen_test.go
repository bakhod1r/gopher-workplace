package chaingen

import "testing"

func TestLookup(t *testing.T) {
	high := map[string]int{"a": 1}
	low := map[string]int{"a": 2, "b": 3}

	if v, ok := Lookup("a", high, low); v != 1 || !ok {
		t.Errorf("Lookup(a) = %v, %v, want 1, true (first layer wins)", v, ok)
	}
	if v, ok := Lookup("b", high, low); v != 3 || !ok {
		t.Errorf("Lookup(b) = %v, %v, want 3, true", v, ok)
	}
	if v, ok := Lookup("missing", high, low); v != 0 || ok {
		t.Errorf("Lookup(missing) = %v, %v, want 0, false", v, ok)
	}
	if _, ok := Lookup[string, int]("a"); ok {
		t.Error("Lookup with no maps reported ok, want false")
	}
}

func TestLookupStoredZeroWins(t *testing.T) {
	high := map[string]int{"a": 0}
	low := map[string]int{"a": 2}
	if v, ok := Lookup("a", high, low); v != 0 || !ok {
		t.Errorf("Lookup(a) = %v, %v, want 0, true (a stored zero is a hit)", v, ok)
	}
}
