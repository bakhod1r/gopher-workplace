package profilepercent

import "testing"

func TestPercent(t *testing.T) {
	cases := []struct {
		value, total int64
		want         float64
	}{
		{1, 3, 33.33},
		{2, 3, 66.67},
		{1, 2, 50},
		{3, 3, 100},
		{0, 3, 0},
		{7, 1000, 0.7},
	}
	for _, c := range cases {
		if got := Percent(c.value, c.total); got != c.want {
			t.Errorf("Percent(%d, %d) = %v, want %v", c.value, c.total, got, c.want)
		}
	}
}

func TestPercentNonPositiveTotal(t *testing.T) {
	for _, total := range []int64{0, -5} {
		if got := Percent(10, total); got != 0 {
			t.Errorf("Percent(10, %d) = %v, want 0", total, got)
		}
	}
}

func TestFormat(t *testing.T) {
	cases := []struct {
		value, total int64
		want         string
	}{
		{1, 3, "33.33%"},
		{3, 3, "100.00%"},
		{1, 2, "50.00%"},
		{10, 0, "0.00%"},
	}
	for _, c := range cases {
		if got := Format(c.value, c.total); got != c.want {
			t.Errorf("Format(%d, %d) = %q, want %q", c.value, c.total, got, c.want)
		}
	}
}
