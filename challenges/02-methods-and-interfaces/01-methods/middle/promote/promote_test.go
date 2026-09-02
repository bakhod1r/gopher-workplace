package promote

import "testing"

func TestPromotedHello(t *testing.T) {
	// Test that Hello is promoted from Base to Extended.
	e := Extended{Base: Base{Name: "Go"}, Extra: "fast"}
	want := "Hello from Go"
	if got := e.Hello(); got != want {
		t.Errorf("Extended.Hello() = %q, want %q", got, want)
	}
}

func TestDescribe(t *testing.T) {
	cases := []struct {
		name string
		e    Extended
		want string
	}{
		{"normal", Extended{Base{"Go"}, "fast"}, "Hello from Go [fast]"},
		{"empty_extra", Extended{Base{"X"}, ""}, "Hello from X []"},
		{"empty_name", Extended{Base{""}, "y"}, "Hello from  [y]"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.e.Describe(); got != tc.want {
				t.Errorf("Describe() = %q, want %q", got, tc.want)
			}
		})
	}
}
