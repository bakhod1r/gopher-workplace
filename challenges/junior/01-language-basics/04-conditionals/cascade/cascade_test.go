package cascade

import "testing"

func TestAccess(t *testing.T) {
	cases := []struct {
		name  string
		level int
		want  string
	}{
		{"admin", 3, "admin,write,read"},
		{"writer", 2, "write,read"},
		{"reader", 1, "read"},
		{"none zero", 0, ""},
		{"none high", 9, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Access(tc.level); got != tc.want {
				t.Errorf("Access(%d) = %q, want %q", tc.level, got, tc.want)
			}
		})
	}
}
