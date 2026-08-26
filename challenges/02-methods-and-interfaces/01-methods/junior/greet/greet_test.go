package greet

import "testing"

func TestGreet(t *testing.T) {
	cases := []struct {
		name string
		p    Person
		want string
	}{
		{"alice", Person{Name: "Alice"}, "Hello, Alice!"},
		{"bob", Person{Name: "Bob"}, "Hello, Bob!"},
		{"empty", Person{Name: ""}, "Hello, !"},
		{"spaces", Person{Name: "Go Gopher"}, "Hello, Go Gopher!"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.p.Greet(); got != tc.want {
				t.Errorf("Person{%q}.Greet() = %q, want %q", tc.p.Name, got, tc.want)
			}
		})
	}
}
