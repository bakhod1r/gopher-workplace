package enumerifc

import "testing"

func TestWords(t *testing.T) {
	w := Words{"a", "b"}
	if got := w.Len(); got != 2 {
		t.Errorf("Len = %d, want 2", got)
	}
	if got := w.At(1); got != "b" {
		t.Errorf("At(1) = %q, want \"b\"", got)
	}
}

func TestJoin(t *testing.T) {
	cases := []struct {
		name string
		w    Words
		sep  string
		want string
	}{
		{"three", Words{"a", "b", "c"}, "-", "a-b-c"},
		{"one", Words{"solo"}, "-", "solo"},
		{"empty", Words{}, "-", ""},
		{"multi_sep", Words{"x", "y"}, ", ", "x, y"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Join(tc.w, tc.sep); got != tc.want {
				t.Errorf("Join = %q, want %q", got, tc.want)
			}
		})
	}
}
