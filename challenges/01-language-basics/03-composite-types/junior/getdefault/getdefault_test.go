package getdefault

import "testing"

func TestGetOr(t *testing.T) {
	m := map[string]int{"a": 1, "zero": 0}
	cases := []struct {
		key  string
		def  int
		want int
	}{
		{"a", 99, 1},
		{"zero", 99, 0}, // present with value 0, not the default
		{"missing", 99, 99},
	}
	for _, c := range cases {
		if got := GetOr(m, c.key, c.def); got != c.want {
			t.Errorf("GetOr(%q,%d)=%d; want %d", c.key, c.def, got, c.want)
		}
	}
	if GetOr(nil, "x", 7) != 7 {
		t.Error("nil map should return default")
	}
}
