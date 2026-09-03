package truncate

import "testing"

func TestTruncateLength(t *testing.T) {
	s := []*Node{{ID: 1}, {ID: 2}, {ID: 3}}
	if got := Truncate(s, 2); len(got) != 2 {
		t.Errorf("len = %d, want 2", len(got))
	}
	if got := Truncate(s, 9); len(got) != 3 {
		t.Errorf("len = %d, want 3 when n exceeds the input", len(got))
	}
	if got := Truncate(s, -1); len(got) != 0 {
		t.Errorf("len = %d, want 0 when n is negative", len(got))
	}
}

func TestTruncateClearsTheTail(t *testing.T) {
	s := []*Node{{ID: 1}, {ID: 2}, {ID: 3}}
	Truncate(s, 1)
	for i := 1; i < len(s); i++ {
		if s[i] != nil {
			t.Errorf("s[%d] still holds %v: the dropped payload stays reachable", i, s[i])
		}
	}
	if s[0] == nil {
		t.Error("s[0] was cleared, but it is kept")
	}
}
