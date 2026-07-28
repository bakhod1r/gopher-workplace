package ringbuffer

import "testing"

func TestAt(t *testing.T) {
	buf := []int{10, 20, 30, 40} // head=2 -> logical order 30,40,10,20
	cases := []struct {
		i, want int
	}{
		{0, 30}, {1, 40}, {2, 10}, {3, 20},
	}
	for _, c := range cases {
		if got := At(buf, 2, c.i); got != c.want {
			t.Errorf("At(head=2,i=%d)=%d; want %d", c.i, got, c.want)
		}
	}
}
