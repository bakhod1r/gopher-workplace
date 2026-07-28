package arrtoslice

import "testing"

func TestAsSlice(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	s := AsSlice(&a)
	if len(s) != 4 {
		t.Fatalf("len=%d want 4", len(s))
	}
	s[3] = 99
	if a[3] != 99 {
		t.Errorf("slice must alias the array")
	}
}
