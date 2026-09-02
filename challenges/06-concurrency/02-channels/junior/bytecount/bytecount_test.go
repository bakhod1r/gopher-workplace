package bytecount

import "testing"

func chanOf(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestTotalBytes(t *testing.T) {
	cases := []struct {
		name  string
		sizes []int
		want  int
	}{
		{"two_responses", []int{1200, 800}, 2000},
		{"single", []int{512}, 512},
		{"empty_window", nil, 0},
		{"empty_bodies", []int{0, 0, 0}, 0},
		{"many", []int{1, 2, 3, 4, 5}, 15},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := TotalBytes(chanOf(tc.sizes...)); got != tc.want {
				t.Errorf("TotalBytes(%v) = %d, want %d", tc.sizes, got, tc.want)
			}
		})
	}
}
