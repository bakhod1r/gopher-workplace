package growcheck

import "testing"

func TestGrewWhenCapacityWasEnough(t *testing.T) {
	s := make([]int, 0, 4)
	if Grew(s, append(s, 1)) {
		t.Error("Grew = true, want false: the capacity was sufficient")
	}
}

func TestGrewWhenCapacityRanOut(t *testing.T) {
	s := make([]int, 1, 1)
	if !Grew(s, append(s, 2)) {
		t.Error("Grew = false, want true: append had to reallocate")
	}
}

func TestGrewFromNil(t *testing.T) {
	var s []int
	if !Grew(s, append(s, 1)) {
		t.Error("Grew = false, want true: a nil slice has no storage")
	}
}

func TestGrewIdentity(t *testing.T) {
	s := make([]int, 2, 4)
	if Grew(s, s) {
		t.Error("Grew(s, s) = true, want false")
	}
	if Grew(s, s[:1]) {
		t.Error("Grew(s, s[:1]) = true, want false: reslicing does not reallocate")
	}
}

func TestGrewAcrossManyAppends(t *testing.T) {
	s := make([]int, 0, 2)
	grew := 0
	for i := 0; i < 64; i++ {
		next := append(s, i)
		if Grew(s, next) {
			grew++
		}
		s = next
	}
	if grew == 0 {
		t.Error("no growth detected across 64 appends, want several")
	}
	if grew > 10 {
		t.Errorf("detected %d reallocations, want the handful append actually performs", grew)
	}
}
