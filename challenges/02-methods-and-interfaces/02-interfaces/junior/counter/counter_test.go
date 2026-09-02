package counter

import "testing"

func TestTotal(t *testing.T) {
	cases := []struct {
		name string
		cs   []Counter
		want int
	}{
		{"mixed", []Counter{&Clicks{N: 3}, Fixed(2)}, 5},
		{"empty", nil, 0},
		{"cancel_out", []Counter{Fixed(-1), Fixed(1)}, 0},
		{"single", []Counter{&Clicks{N: 42}}, 42},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Total(tc.cs); got != tc.want {
				t.Errorf("Total = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	if got := (&Clicks{N: 7}).Count(); got != 7 {
		t.Errorf("Clicks.Count = %d, want 7", got)
	}
	if got := Fixed(9).Count(); got != 9 {
		t.Errorf("Fixed.Count = %d, want 9", got)
	}
}
