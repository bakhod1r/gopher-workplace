package countbygen

import "testing"

func TestCountBy(t *testing.T) {
	parity := func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	}
	got := CountBy([]int{1, 2, 3}, parity)
	if got["odd"] != 2 || got["even"] != 1 {
		t.Errorf("CountBy = %v, want {odd:2 even:1}", got)
	}
	if len(got) != 2 {
		t.Errorf("len = %d, want 2", len(got))
	}
}

func TestCountByEmpty(t *testing.T) {
	got := CountBy([]int{}, func(n int) int { return n })
	if got == nil {
		t.Fatal("CountBy(empty) = nil, want an empty non-nil map")
	}
	if len(got) != 0 {
		t.Errorf("CountBy(empty) = %v, want {}", got)
	}
}

func TestCountByStructs(t *testing.T) {
	type req struct{ status int }
	got := CountBy([]req{{200}, {200}, {500}}, func(r req) int { return r.status / 100 })
	if got[2] != 2 || got[5] != 1 {
		t.Errorf("CountBy = %v, want {2:2 5:1}", got)
	}
}
