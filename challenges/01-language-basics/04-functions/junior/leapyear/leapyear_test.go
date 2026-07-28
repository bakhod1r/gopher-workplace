package leapyear

import "testing"

func TestIsLeap(t *testing.T) {
	cases := []struct {
		y    int
		want bool
	}{
		{2000, true}, {1900, false}, {2024, true}, {2023, false}, {2100, false},
	}
	for _, c := range cases {
		if got := IsLeap(c.y); got != c.want {
			t.Errorf("IsLeap(%d)=%v want %v", c.y, got, c.want)
		}
	}
}
