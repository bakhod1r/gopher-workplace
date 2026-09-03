package fieldpad

import "testing"

func TestAlignUp(t *testing.T) {
	cases := []struct{ n, a, want int }{
		{9, 8, 16},
		{8, 8, 8},
		{0, 8, 0},
		{1, 2, 2},
		{7, 1, 7},
		{7, 0, 7},
		{15, 4, 16},
	}
	for _, c := range cases {
		if got := AlignUp(c.n, c.a); got != c.want {
			t.Errorf("AlignUp(%d, %d) = %d, want %d", c.n, c.a, got, c.want)
		}
	}
}

func TestStructSize(t *testing.T) {
	cases := []struct {
		sizes []int
		want  int
	}{
		{[]int{8, 4, 2, 1}, 16}, // widest first: 0,8,12,14 -> 15 -> 16
		{[]int{1, 8, 2, 4}, 24}, // bool first: 0,8,16,20 -> 24
		{[]int{8}, 8},
		{[]int{1, 1, 1}, 3},
		{nil, 0},
		{[]int{4, 4}, 8},
		{[]int{2, 8, 2}, 24}, // 0, 8, 16 -> 18 -> 24
	}
	for _, c := range cases {
		if got := StructSize(c.sizes); got != c.want {
			t.Errorf("StructSize(%v) = %d, want %d", c.sizes, got, c.want)
		}
	}
}

func TestStructSizeSkipsZeroFields(t *testing.T) {
	if got := StructSize([]int{8, 0, 4}); got != 16 {
		t.Errorf("StructSize = %d, want 16", got)
	}
}
