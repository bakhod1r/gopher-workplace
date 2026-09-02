package sayhello

import "testing"

func TestHello(t *testing.T) {
	if got := (English{}).Hello(); got != "Hello" {
		t.Errorf("English.Hello = %q", got)
	}
	if got := (Uzbek{}).Hello(); got != "Salom" {
		t.Errorf("Uzbek.Hello = %q", got)
	}
}

func TestGreet(t *testing.T) {
	cases := []struct {
		name string
		g    Greeter
		who  string
		want string
	}{
		{"english", English{}, "Ann", "Hello, Ann"},
		{"uzbek", Uzbek{}, "Ali", "Salom, Ali"},
		{"empty_name", English{}, "", "Hello, "},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Greet(tc.g, tc.who); got != tc.want {
				t.Errorf("Greet = %q, want %q", got, tc.want)
			}
		})
	}
}
