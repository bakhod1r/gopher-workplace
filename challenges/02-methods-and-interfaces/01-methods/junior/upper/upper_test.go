package upper

import "testing"

func TestUpper(t *testing.T) {
	cases := []struct {
		name string
		s    MyString
		want string
	}{
		{"lower", "hello", "HELLO"},
		{"mixed", "Go Gopher", "GO GOPHER"},
		{"already", "ABC", "ABC"},
		{"empty", "", ""},
		{"numbers", "go123", "GO123"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.s.Upper(); got != tc.want {
				t.Errorf("MyString(%q).Upper() = %q, want %q", string(tc.s), got, tc.want)
			}
		})
	}
}
