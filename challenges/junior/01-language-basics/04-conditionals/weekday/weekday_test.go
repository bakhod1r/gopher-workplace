package weekday

import "testing"

func TestWeekday(t *testing.T) {
	cases := []struct {
		name string
		d    int
		want string
	}{
		{"monday", 1, "Monday"},
		{"wednesday", 3, "Wednesday"},
		{"sunday", 7, "Sunday"},
		{"zero", 0, "Unknown"},
		{"too big", 8, "Unknown"},
		{"negative", -1, "Unknown"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Weekday(tc.d); got != tc.want {
				t.Errorf("Weekday(%d) = %q, want %q", tc.d, got, tc.want)
			}
		})
	}
}
