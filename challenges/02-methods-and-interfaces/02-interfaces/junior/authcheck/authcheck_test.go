package authcheck

import "testing"

func TestCanEnter(t *testing.T) {
	cases := []struct {
		name string
		a    Authenticator
		want bool
	}{
		{"valid_token", Token{Value: "abc"}, true},
		{"expired_token", Token{Value: "expired"}, false},
		{"empty_token", Token{Value: ""}, false},
		{"guest", Guest{}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CanEnter(tc.a); got != tc.want {
				t.Errorf("CanEnter(%#v) = %v, want %v", tc.a, got, tc.want)
			}
		})
	}
}
