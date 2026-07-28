package commaok

import "testing"

func TestLookup(t *testing.T) {
	scores := map[string]int{"a": 0, "b": 5}
	if s, ok := Lookup(scores, "a"); s != 0 || !ok {
		t.Errorf("a=%d,%v want 0,true", s, ok)
	}
	if s, ok := Lookup(scores, "z"); s != 0 || ok {
		t.Errorf("z=%d,%v want 0,false", s, ok)
	}
}
