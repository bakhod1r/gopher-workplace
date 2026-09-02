package resizeworker

import "testing"

func TestScaleRequest(t *testing.T) {
	cases := []struct {
		name  string
		width int
		want  int
	}{
		{"typical", 640, 1280},
		{"zero", 0, 0},
		{"negative", -3, -6},
		{"one", 1, 2},
		{"large", 1920, 3840},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ScaleRequest(tc.width); got != tc.want {
				t.Errorf("ScaleRequest(%d) = %d, want %d", tc.width, got, tc.want)
			}
		})
	}
}
