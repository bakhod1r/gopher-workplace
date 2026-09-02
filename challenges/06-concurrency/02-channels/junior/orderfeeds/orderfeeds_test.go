package orderfeeds

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

func TestMergeFeeds(t *testing.T) {
	cases := []struct {
		name             string
		primary, standby []int
		want             []int
	}{
		{"two_then_one", []int{1, 2}, []int{3}, []int{1, 2, 3}},
		{"primary_empty", nil, []int{9}, []int{9}},
		{"standby_empty", []int{4}, nil, []int{4}},
		{"both_empty", nil, nil, []int{}},
		{"order_preserved", []int{5, 1}, []int{7, 2}, []int{5, 1, 7, 2}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := MergeFeeds(chanOf(tc.primary...), chanOf(tc.standby...))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MergeFeeds(%v, %v) = %#v, want %#v",
					tc.primary, tc.standby, got, tc.want)
			}
		})
	}
}
