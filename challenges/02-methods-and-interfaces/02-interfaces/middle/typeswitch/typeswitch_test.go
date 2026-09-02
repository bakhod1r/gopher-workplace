package typeswitch

import "testing"

func TestRender(t *testing.T) {
	cases := []struct {
		name string
		in   any
		want string
	}{
		{"int", 42, "42"},
		{"negative_int", -7, "-7"},
		{"bool_true", true, "true"},
		{"bool_false", false, "false"},
		{"string", "hi", "hi"},
		{"empty_string", "", ""},
		{"slice", []string{"a", "b"}, "a,b"},
		{"empty_slice", []string{}, ""},
		{"float", 3.5, "?"},
		{"nil", nil, "?"},
		{"other_slice", []int{1}, "?"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Render(tc.in); got != tc.want {
				t.Errorf("Render(%#v) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
