package runecount

import "testing"

func TestCount(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want int
	}{
		{"ascii", "abc", 3},
		{"accented", "héllo", 5},
		{"cjk", "日本", 2},
		{"empty", "", 0},
		{"emoji", "a🙂b", 3},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Count(tc.in); got != tc.want {
				t.Errorf("Count(%q) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}
