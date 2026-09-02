package ptrofgen

import "testing"

func TestPtr(t *testing.T) {
	p := Ptr(7)
	if p == nil {
		t.Fatal("Ptr(7) = nil, want a non-nil pointer")
	}
	if *p != 7 {
		t.Errorf("*Ptr(7) = %v, want 7", *p)
	}
	s := Ptr("go")
	if *s != "go" {
		t.Errorf("*Ptr(%q) = %q, want %q", "go", *s, "go")
	}
	if a, b := Ptr(1), Ptr(1); a == b {
		t.Error("Ptr(1) and Ptr(1) returned the same pointer, want distinct allocations")
	}
}
