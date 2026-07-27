package ifinit

import "testing"

func TestBucket(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want string
	}{
		{"nine", 9, "zero"},
		{"ten", 10, "one"},
		{"eleven", 11, "two"},
		{"zero", 0, "zero"},
		{"one", 1, "one"},
		{"two", 2, "two"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Bucket(tc.n); got != tc.want {
				t.Errorf("Bucket(%d) = %q, want %q", tc.n, got, tc.want)
			}
		})
	}
}
