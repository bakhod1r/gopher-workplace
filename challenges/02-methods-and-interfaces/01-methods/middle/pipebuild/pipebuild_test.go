package pipebuild

import "testing"

func TestPipe(t *testing.T) {
	cases := []struct {
		name string
		pipe *Pipe
		want string
	}{
		{"upper", NewPipe("hello").Upper(), "HELLO"},
		{"replace", NewPipe("a b a").Replace("a", "x"), "x b x"},
		{"chain", NewPipe("go lang").Upper().Replace(" ", "-"), "GO-LANG"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.pipe.Result(); got != tc.want {
				t.Errorf("Result() = %q, want %q", got, tc.want)
			}
		})
	}
}
