package nilslicedata

import "testing"

func TestFirstOr(t *testing.T) {
	if FirstOr(nil, 7) != 7 {
		t.Errorf("empty should give default (no panic)")
	}
	if FirstOr([]int{5}, 7) != 5 {
		t.Errorf("want 5")
	}
}
