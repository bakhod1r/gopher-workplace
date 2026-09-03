package metricsflush

import (
	"reflect"
	"testing"
)

func collect(in <-chan int, size int) [][]int {
	var got [][]int
	for b := range FlushBatches(in, size) {
		got = append(got, b)
	}
	return got
}

func feed(samples ...int) <-chan int {
	ch := make(chan int, len(samples))
	for _, s := range samples {
		ch <- s
	}
	close(ch)
	return ch
}

func TestFlushBatches(t *testing.T) {
	cases := []struct {
		name    string
		samples []int
		size    int
		want    [][]int
	}{
		{"exact_fit", []int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
		{"trailing_partial", []int{1, 2, 3}, 2, [][]int{{1, 2}, {3}}},
		{"one_batch", []int{1, 2}, 5, [][]int{{1, 2}}},
		{"empty_input", nil, 3, nil},
		{"size_one", []int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
		{"size_zero_means_one", []int{7, 8}, 0, [][]int{{7}, {8}}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collect(feed(tc.samples...), tc.size)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("batches = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestOutputChannelIsClosed(t *testing.T) {
	out := FlushBatches(feed(1, 2, 3), 2)
	for range out {
	}
	if _, ok := <-out; ok {
		t.Error("output channel was not closed after the input drained")
	}
}

func TestBatchesAreIndependentSlices(t *testing.T) {
	out := FlushBatches(feed(1, 2, 3, 4), 2)
	first := <-out
	second := <-out
	first[0] = 99
	if second[0] != 3 {
		t.Errorf("batches share a backing array: second = %v", second)
	}
}

func TestStreamingInput(t *testing.T) {
	in := make(chan int)
	out := FlushBatches(in, 2)

	go func() {
		for i := 1; i <= 5; i++ {
			in <- i
		}
		close(in)
	}()

	var got [][]int
	for b := range out {
		got = append(got, b)
	}
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("batches = %v, want %v", got, want)
	}
}
