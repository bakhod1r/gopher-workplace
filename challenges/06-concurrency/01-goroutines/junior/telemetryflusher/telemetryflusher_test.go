package telemetryflusher

import (
	"reflect"
	"testing"
)

func TestBatchErrorCounts(t *testing.T) {
	cases := []struct {
		name  string
		codes []int
		batch int
		want  []int
	}{
		{"even_batches", []int{200, 500, 503, 200}, 2, []int{1, 1}},
		{"ragged_tail", []int{500, 500, 200}, 2, []int{2, 0}},
		{"one_per_batch", []int{500, 200}, 1, []int{1, 0}},
		{"single_batch", []int{404, 500, 502}, 10, []int{2}},
		{"bad_size", []int{200}, 0, []int(nil)},
		{"empty", []int{}, 3, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := BatchErrorCounts(tc.codes, tc.batch); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("BatchErrorCounts(%v) = %v, want %v", tc.codes, got, tc.want)
			}
		})
	}
}
