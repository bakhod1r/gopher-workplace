package sayhello

import "testing"

func TestSayHello(t *testing.T) {
	cases := []struct {
		name string
		g    Greeter
		want string
	}{
		{"english", English{}, "Hello!"},
		{"uzbek", Uzbek{}, "Salom!"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SayHello(tc.g); got != tc.want {
				t.Errorf("SayHello() = %q, want %q", got, tc.want)
			}
		})
	}
}
