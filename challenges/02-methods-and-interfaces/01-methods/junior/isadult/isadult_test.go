package isadult

import "testing"

func TestIsAdult(t *testing.T) {
	cases := []struct {
		name string
		p    Person
		want bool
	}{
		{"25", Person{"Alice", 25}, true},
		{"17", Person{"Bob", 17}, false},
		{"exactly_18", Person{"Carol", 18}, true},
		{"zero", Person{"Baby", 0}, false},
		{"negative", Person{"Bug", -1}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.p.IsAdult(); got != tc.want {
				t.Errorf("Person{%q,%d}.IsAdult() = %v, want %v",
					tc.p.Name, tc.p.Age, got, tc.want)
			}
		})
	}
}
