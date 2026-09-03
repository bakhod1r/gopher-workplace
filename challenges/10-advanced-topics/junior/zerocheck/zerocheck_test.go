package zerocheck

import "testing"

type pair struct {
	A int
	B string
}

func TestIsZero(t *testing.T) {
	cases := []struct {
		in   any
		want bool
	}{
		{nil, true},
		{0, true},
		{1, false},
		{"", true},
		{"x", false},
		{false, true},
		{true, false},
		{pair{}, true},
		{pair{A: 1}, false},
		{(*pair)(nil), true},
		{&pair{}, false},
		{[]int(nil), true},
		{[]int{}, false},
		{0.0, true},
	}
	for _, c := range cases {
		if got := IsZero(c.in); got != c.want {
			t.Errorf("IsZero(%#v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestIsZeroArrays(t *testing.T) {
	if !IsZero([2]int{}) {
		t.Error("IsZero([2]int{}) = false, want true")
	}
	if IsZero([2]int{0, 1}) {
		t.Error("IsZero([2]int{0,1}) = true, want false")
	}
}
