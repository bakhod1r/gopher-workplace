package keystringgen

import "testing"

func TestNormalize(t *testing.T) {
	got := Normalize(map[HeaderName]int{"Content-Type": 1, "ACCEPT": 2})
	if got["content-type"] != 1 || got["accept"] != 2 {
		t.Errorf("Normalize = %v, want lowercase keys", got)
	}
	if len(got) != 2 {
		t.Errorf("len = %d, want 2", len(got))
	}
}

func TestNormalizeCollisions(t *testing.T) {
	got := Normalize(map[HeaderName]int{"A": 1, "a": 1})
	if len(got) != 1 {
		t.Errorf("len = %d, want 1 (keys collapse after lowercasing)", len(got))
	}
}

func TestNormalizeEmpty(t *testing.T) {
	got := Normalize(map[HeaderName]int(nil))
	if got == nil {
		t.Fatal("Normalize(nil) = nil, want an empty non-nil map")
	}
	if len(got) != 0 {
		t.Errorf("Normalize(nil) = %v, want {}", got)
	}
}
