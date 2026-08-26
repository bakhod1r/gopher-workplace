package fullname

import "testing"

func TestFullName(t *testing.T) {
	cases := []struct {
		name string
		p    Person
		want string
	}{
		{"normal", Person{"Go", "Gopher"}, "Go Gopher"},
		{"empty_first", Person{"", "Doe"}, " Doe"},
		{"empty_last", Person{"Jane", ""}, "Jane "},
		{"both_empty", Person{"", ""}, " "},
		{"unicode", Person{"Ïslom", "Karimov"}, "Ïslom Karimov"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.p.FullName(); got != tc.want {
				t.Errorf("Person{%q,%q}.FullName() = %q, want %q",
					tc.p.First, tc.p.Last, got, tc.want)
			}
		})
	}
}
