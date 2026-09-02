package cachewarm

import "testing"

func chanOf(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestNextKey(t *testing.T) {
	cases := []struct {
		name   string
		feed   []int
		wantID int
		wantOK bool
	}{
		{"key_available", []int{5}, 5, true},
		{"feed_finished", nil, 0, false},
		{"key_zero", []int{0}, 0, true},
		{"first_of_many", []int{7, 8, 9}, 7, true},
		{"negative_id", []int{-4}, -4, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotID, gotOK := NextKey(chanOf(tc.feed...))
			if gotID != tc.wantID || gotOK != tc.wantOK {
				t.Errorf("NextKey(%v) = %d, %t, want %d, %t",
					tc.feed, gotID, gotOK, tc.wantID, tc.wantOK)
			}
		})
	}
}
