package stringer

import (
	"fmt"
	"testing"
)

func TestStringer(t *testing.T) {
	cases := []struct {
		c    Color
		want string
	}{
		{Color{255, 0, 128}, "#ff0080"},
		{Color{0, 0, 0}, "#000000"},
		{Color{255, 255, 255}, "#ffffff"},
	}

	for _, tc := range cases {
		got := fmt.Sprint(tc.c)
		if got != tc.want {
			t.Errorf("Color%v = %q, want %q", tc.c, got, tc.want)
		}
	}
}
