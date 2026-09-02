package bytestringgen

import "testing"

func TestHasPrefix(t *testing.T) {
	if !HasPrefix("GET /x", "GET") {
		t.Error(`HasPrefix("GET /x", "GET") = false, want true`)
	}
	if HasPrefix([]byte("POST /x"), "GET") {
		t.Error(`HasPrefix([]byte("POST /x"), "GET") = true, want false`)
	}
	if !HasPrefix([]byte("GET /x"), "GET") {
		t.Error(`HasPrefix([]byte("GET /x"), "GET") = false, want true`)
	}
	if !HasPrefix("anything", "") {
		t.Error("an empty prefix should always match")
	}
}

func TestSize(t *testing.T) {
	if got := Size("abc"); got != 3 {
		t.Errorf(`Size("abc") = %d, want 3`, got)
	}
	if got := Size([]byte{1, 2}); got != 2 {
		t.Errorf("Size([]byte{1, 2}) = %d, want 2", got)
	}
	if got := Size([]byte(nil)); got != 0 {
		t.Errorf("Size(nil) = %d, want 0", got)
	}
}
