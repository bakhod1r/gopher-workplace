package ordinal

import "testing"

func TestFormat(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{
		{1, "1st"}, {2, "2nd"}, {3, "3rd"}, {4, "4th"},
		{11, "11th"}, {12, "12th"}, {13, "13th"}, // teens are all "th"
		{21, "21st"}, {112, "112th"}, {101, "101st"},
	}
	for _, c := range cases {
		if got := Format(c.n); got != c.want {
			t.Errorf("Format(%d)=%q; want %q", c.n, got, c.want)
		}
	}
}
