package nilout

import "testing"

func TestDropAllNilsEveryElement(t *testing.T) {
	s := []*Node{{ID: 1}, {ID: 2}, {ID: 3}}
	DropAll(s)
	if len(s) != 3 {
		t.Fatalf("len = %d, want 3: the length must not change", len(s))
	}
	for i, p := range s {
		if p != nil {
			t.Errorf("s[%d] = %v, want nil", i, p)
		}
	}
}

func TestDropAllIsVisibleToTheCaller(t *testing.T) {
	s := []*Node{{ID: 1}}
	DropAll(s[:1])
	if s[0] != nil {
		t.Error("the caller's element was not cleared")
	}
}

func TestDropAllEmpty(t *testing.T) {
	DropAll(nil)
}
