package capgrew

import "testing"

func TestAppendedValue(t *testing.T) {
	got, _ := Appended([]int{1, 2}, 3)
	if len(got) != 3 || got[2] != 3 {
		t.Errorf("Appended = %v, want [1 2 3]", got)
	}
}

func TestAppendedWithRoom(t *testing.T) {
	s := make([]int, 0, 4)
	got, grew := Appended(s, 1)
	if grew {
		t.Error("grew = true, want false: the capacity was sufficient")
	}
	if len(got) != 1 || cap(got) != 4 {
		t.Errorf("len, cap = %d, %d, want 1, 4", len(got), cap(got))
	}
}

func TestAppendedWithoutRoom(t *testing.T) {
	s := make([]int, 1, 1)
	got, grew := Appended(s, 2)
	if !grew {
		t.Error("grew = false, want true: the capacity was exhausted")
	}
	if cap(got) <= 1 {
		t.Errorf("cap = %d, want more than 1", cap(got))
	}
}

func TestAppendedFromNil(t *testing.T) {
	got, grew := Appended(nil, 5)
	if !grew {
		t.Error("grew = false, want true: a nil slice has no capacity")
	}
	if len(got) != 1 || got[0] != 5 {
		t.Errorf("Appended = %v, want [5]", got)
	}
}
