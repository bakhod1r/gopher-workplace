package anyslicepitfall

import "testing"

func TestToAny(t *testing.T) {
	got := ToAny([]int{1, 2})
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	if got[0] != any(1) || got[1] != any(2) {
		t.Errorf("ToAny = %v, want [1 2]", got)
	}
	strs := ToAny([]string{"a"})
	if len(strs) != 1 || strs[0] != any("a") {
		t.Errorf("ToAny = %v, want [a]", strs)
	}
	empty := ToAny([]int{})
	if empty == nil || len(empty) != 0 {
		t.Errorf("ToAny([]) = %v, want an empty non-nil slice", empty)
	}
}

func TestSumIntsNeedsNoConversion(t *testing.T) {
	if got := SumInts([]int{1, 2, 3}); got != 6 {
		t.Errorf("SumInts = %v, want 6", got)
	}
}
