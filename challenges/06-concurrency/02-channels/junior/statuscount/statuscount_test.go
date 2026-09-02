package statuscount

import "testing"

func chanOf(vals ...string) <-chan string {
	ch := make(chan string, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestCountStatus(t *testing.T) {
	cases := []struct {
		name  string
		lines []string
		want  string
		n     int
	}{
		{"two_matches", []string{"200", "500", "200"}, "200", 2},
		{"no_match", []string{"200"}, "404", 0},
		{"empty_stream", nil, "200", 0},
		{"all_match", []string{"503", "503", "503"}, "503", 3},
		{"case_sensitive", []string{"OK", "ok"}, "ok", 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountStatus(chanOf(tc.lines...), tc.want); got != tc.n {
				t.Errorf("CountStatus(%v, %q) = %d, want %d",
					tc.lines, tc.want, got, tc.n)
			}
		})
	}
}
