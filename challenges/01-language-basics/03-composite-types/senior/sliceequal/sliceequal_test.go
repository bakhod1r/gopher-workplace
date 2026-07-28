package sliceequal

import "testing"

func TestEqual(t *testing.T) {
	cases := []struct {
		a, b []int
		want bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{1, 9, 3}, false},
		{[]int{1, 2}, []int{1, 2, 3}, false},
		{nil, nil, true},
	}
	for _, c := range cases {
		if got := Equal(c.a, c.b); got != c.want {
			t.Errorf("Equal(%v,%v)=%v; want %v", c.a, c.b, got, c.want)
		}
	}
}
