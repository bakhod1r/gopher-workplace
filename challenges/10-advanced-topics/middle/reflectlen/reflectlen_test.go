package reflectlen

import "testing"

func TestLength(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	cases := []struct {
		in any
		n  int
		ok bool
	}{
		{[]int{1, 2}, 2, true},
		{"héllo", 6, true},
		{[3]int{}, 3, true},
		{map[string]int{"a": 1, "b": 2}, 2, true},
		{ch, 1, true},
		{[]int(nil), 0, true},
		{map[string]int(nil), 0, true},
		{"", 0, true},
		{3, 0, false},
		{nil, 0, false},
		{struct{}{}, 0, false},
		{&[3]int{}, 0, false},
	}
	for _, c := range cases {
		n, ok := Length(c.in)
		if n != c.n || ok != c.ok {
			t.Errorf("Length(%#v) = %d, %v, want %d, %v", c.in, n, ok, c.n, c.ok)
		}
	}
}

func TestLengthCountsBytesNotRunes(t *testing.T) {
	if n, _ := Length("日本"); n != 6 {
		t.Errorf("Length = %d, want 6: a string's length is in bytes", n)
	}
}

func TestLengthDoesNotPanic(t *testing.T) {
	for _, in := range []any{nil, 3, 1.5, true, func() {}, struct{ A int }{}} {
		_, _ = Length(in)
	}
}
