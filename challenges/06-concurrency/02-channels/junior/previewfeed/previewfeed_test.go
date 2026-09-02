package previewfeed

import (
	"reflect"
	"testing"
)

func chanOf(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestPreviewOrders(t *testing.T) {
	cases := []struct {
		name  string
		feed  []int
		limit int
		want  []int
	}{
		{"first_two", []int{1, 2, 3}, 2, []int{1, 2}},
		{"feed_shorter_than_limit", []int{1}, 5, []int{1}},
		{"no_preview", []int{1, 2}, 0, []int{}},
		{"exact_limit", []int{4, 5}, 2, []int{4, 5}},
		{"empty_feed", nil, 3, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := PreviewOrders(chanOf(tc.feed...), tc.limit)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PreviewOrders(%v, %d) = %#v, want %#v",
					tc.feed, tc.limit, got, tc.want)
			}
		})
	}
}
