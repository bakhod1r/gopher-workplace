package luhn

import "testing"

func TestValid(t *testing.T) {
	cases := []struct {
		s  string
		ok bool
	}{
		{"4539148803436467", true}, // valid test card
		{"8273123273520569", false},
		{"79927398713", true},
		{"79927398710", false},
		{"1234a", false},
		{"", false},
	}
	for _, c := range cases {
		if got := Valid(c.s); got != c.ok {
			t.Errorf("Valid(%q)=%v; want %v", c.s, got, c.ok)
		}
	}
}
