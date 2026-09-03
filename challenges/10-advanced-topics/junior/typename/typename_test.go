package typename

import "testing"

type widget struct{ A int }

func TestTypeName(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{3, "int"},
		{"s", "string"},
		{3.5, "float64"},
		{[]int{1}, "[]int"},
		{map[string]int{}, "map[string]int"},
		{widget{}, "typename.widget"},
		{&widget{}, "*typename.widget"},
		{nil, "<nil>"},
		{[]byte("x"), "[]uint8"},
	}
	for _, c := range cases {
		if got := TypeName(c.in); got != c.want {
			t.Errorf("TypeName(%#v) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestTypeNameNamedTypes(t *testing.T) {
	type alias = int
	type named int
	if got := TypeName(alias(1)); got != "int" {
		t.Errorf("TypeName = %q, want \"int\": an alias is not a new type", got)
	}
	if got := TypeName(named(1)); got != "typename.named" {
		t.Errorf("TypeName = %q, want \"typename.named\"", got)
	}
}
