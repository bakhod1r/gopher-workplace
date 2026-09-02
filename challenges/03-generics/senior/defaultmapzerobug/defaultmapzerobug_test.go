package defaultmapzerobug

import "testing"

func TestGetReturnsDefault(t *testing.T) {
	d := &DefaultMap[string, int]{Default: 42}
	if got := d.Get("nope"); got != 42 {
		t.Errorf("Get = %d, want 42", got)
	}
}

func TestGetDoesNotInsert(t *testing.T) {
	d := &DefaultMap[string, int]{Default: 42}
	d.Get("a")
	d.Get("b")
	if got := d.Len(); got != 0 {
		t.Errorf("Len = %d after reads only, want 0", got)
	}
}

func TestSetThenGet(t *testing.T) {
	d := &DefaultMap[string, int]{Default: 42}
	d.Set("a", 1)
	if got := d.Get("a"); got != 1 {
		t.Errorf("Get = %d, want 1", got)
	}
	if got := d.Len(); got != 1 {
		t.Errorf("Len = %d, want 1", got)
	}
}
