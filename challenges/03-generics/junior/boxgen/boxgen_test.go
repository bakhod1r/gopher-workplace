package boxgen

import "testing"

func TestBox(t *testing.T) {
	var b Box[int]
	if v, ok := b.Get(); v != 0 || ok {
		t.Errorf("Get() on empty box = %v, %v, want 0, false", v, ok)
	}
	b.Set(5)
	if v, ok := b.Get(); v != 5 || !ok {
		t.Errorf("Get() = %v, %v, want 5, true", v, ok)
	}
}

func TestBoxStoredZeroCounts(t *testing.T) {
	var b Box[int]
	b.Set(0)
	if v, ok := b.Get(); v != 0 || !ok {
		t.Errorf("Get() = %v, %v, want 0, true (a stored zero is still stored)", v, ok)
	}
}

func TestBoxStrings(t *testing.T) {
	var b Box[string]
	b.Set("")
	if _, ok := b.Get(); !ok {
		t.Error(`Get() after Set("") reported false, want true`)
	}
}
