package kindof

import "testing"

type point struct{ X int }

func TestKindName(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{3, "int"},
		{"s", "string"},
		{3.5, "float64"},
		{[]int{1}, "slice"},
		{map[string]int{}, "map"},
		{point{}, "struct"},
		{&point{}, "ptr"},
		{true, "bool"},
		{nil, "invalid"},
	}
	for _, c := range cases {
		if got := KindName(c.in); got != c.want {
			t.Errorf("KindName(%#v) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestKindNameIgnoresTheNamedType(t *testing.T) {
	type myInt int
	if got := KindName(myInt(1)); got != "int" {
		t.Errorf("KindName = %q, want \"int\": a kind is not a type name", got)
	}
}
