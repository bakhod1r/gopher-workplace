package grade

import "testing"

func TestGrade(t *testing.T) {
	cases := []struct {
		name  string
		score int
		want  string
	}{
		{"top", 100, "A"},
		{"A boundary", 90, "A"},
		{"B", 85, "B"},
		{"C boundary", 70, "C"},
		{"D", 65, "D"},
		{"F just below", 59, "F"},
		{"F zero", 0, "F"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Grade(tc.score); got != tc.want {
				t.Errorf("Grade(%d) = %q, want %q", tc.score, got, tc.want)
			}
		})
	}
}
