package tagless

import "testing"

func TestClassify(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want string
	}{
		{"negative", -5, "negative"},
		{"zero", 0, "zero"},
		{"positive", 7, "positive"},
		{"minus one", -1, "negative"},
		{"one", 1, "positive"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Classify(tc.n); got != tc.want {
				t.Errorf("Classify(%d) = %q, want %q", tc.n, got, tc.want)
			}
		})
	}
}
