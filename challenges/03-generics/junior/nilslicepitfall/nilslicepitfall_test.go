package nilslicepitfall

import "testing"

func TestCollectNeverNil(t *testing.T) {
	none := func(int) bool { return false }
	got := Collect([]int{1, 2, 3}, none)
	if got == nil {
		t.Fatal("Collect returned nil, want an empty non-nil slice")
	}
	if len(got) != 0 {
		t.Errorf("Collect = %v, want []", got)
	}
	if IsNil(got) {
		t.Error("IsNil(Collect(...)) = true, want false")
	}
}

func TestCollectKeeps(t *testing.T) {
	even := func(n int) bool { return n%2 == 0 }
	got := Collect([]int{1, 2, 3, 4}, even)
	if len(got) != 2 || got[0] != 2 || got[1] != 4 {
		t.Errorf("Collect = %v, want [2 4]", got)
	}
}

func TestIsNil(t *testing.T) {
	if !IsNil([]int(nil)) {
		t.Error("IsNil(nil) = false, want true")
	}
	if IsNil([]int{}) {
		t.Error("IsNil([]int{}) = true, want false")
	}
}
