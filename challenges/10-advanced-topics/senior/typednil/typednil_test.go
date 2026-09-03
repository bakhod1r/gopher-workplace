package typednil

import "testing"

type myErr struct{}

func (*myErr) Error() string { return "boom" }

func TestIsNilValue(t *testing.T) {
	var p *int
	var m map[string]int
	var s []int
	var c chan int
	var f func()
	cases := []struct {
		name string
		in   any
		want bool
	}{
		{"untyped nil", nil, true},
		{"nil pointer", p, true},
		{"nil map", m, true},
		{"nil slice", s, true},
		{"nil chan", c, true},
		{"nil func", f, true},
		{"non-nil pointer", new(int), false},
		{"empty slice", []int{}, false},
		{"zero int", 0, false},
		{"empty string", "", false},
		{"struct", struct{}{}, false},
	}
	for _, c := range cases {
		if got := IsNilValue(c.in); got != c.want {
			t.Errorf("%s: IsNilValue = %v, want %v", c.name, got, c.want)
		}
	}
}

func TestIsNilValueCatchesTheTypedNilError(t *testing.T) {
	var e error = (*myErr)(nil)
	if e == nil {
		t.Fatal("the fixture is wrong: a typed nil error is not == nil")
	}
	if !IsNilValue(e) {
		t.Error("IsNilValue = false for a typed nil error: this is the case the function is for")
	}
}

func TestIsNilValueDoesNotPanic(t *testing.T) {
	for _, in := range []any{0, "", 1.5, [2]int{}, struct{ A int }{}} {
		_ = IsNilValue(in)
	}
}
