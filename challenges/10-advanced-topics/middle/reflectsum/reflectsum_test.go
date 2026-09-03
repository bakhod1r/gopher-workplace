package reflectsum

import "testing"

type myInt int

func TestSum(t *testing.T) {
	cases := []struct {
		in    any
		total int64
		ok    bool
	}{
		{[]int{1, 2, 3}, 6, true},
		{[]int32{-1, 1}, 0, true},
		{[]int64{1 << 40, 1 << 40}, 1 << 41, true},
		{[3]int{1, 2, 3}, 6, true},
		{[]myInt{2, 3}, 5, true},
		{[]int{}, 0, true},
		{[]string{"a"}, 0, false},
		{[]uint8{1}, 0, false},
		{3, 0, false},
		{nil, 0, false},
		{map[string]int{"a": 1}, 0, false},
	}
	for _, c := range cases {
		total, ok := Sum(c.in)
		if total != c.total || ok != c.ok {
			t.Errorf("Sum(%#v) = %d, %v, want %d, %v", c.in, total, ok, c.total, c.ok)
		}
	}
}

func TestSumWideAccumulator(t *testing.T) {
	in := make([]int32, 8)
	for i := range in {
		in[i] = 1 << 30
	}
	if total, ok := Sum(in); !ok || total != 8<<30 {
		t.Errorf("Sum = %d, %v, want %d, true: the total must not overflow", total, ok, int64(8)<<30)
	}
}
