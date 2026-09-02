package ordefault

import "testing"

func TestOrDefault(t *testing.T) {
	cases := []struct {
		name string
		v    string
		err  error
		def  string
		want string
	}{
		{"ok", "8080", nil, "80", "8080"},
		{"ok_empty", "", nil, "80", ""},
		{"failed", "", ErrMissing, "80", "80"},
		{"failed_with_value", "junk", ErrMissing, "80", "80"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := OrDefault(tc.v, tc.err, tc.def); got != tc.want {
				t.Errorf("OrDefault(%q, %v, %q) = %q, want %q", tc.v, tc.err, tc.def, got, tc.want)
			}
		})
	}
}
