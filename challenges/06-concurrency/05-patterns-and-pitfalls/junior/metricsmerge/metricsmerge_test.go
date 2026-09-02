package metricsmerge

import "testing"

func send(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestMergeMetrics(t *testing.T) {
	cases := []struct {
		name  string
		nodes [][]int
		want  []int
	}{
		{"two_nodes", [][]int{{1}, {2}}, []int{1, 2}},
		{"uneven_nodes", [][]int{{3, 1}, {2}}, []int{1, 2, 3}},
		{"single_node", [][]int{{5, 4}}, []int{4, 5}},
		{"silent_nodes", [][]int{{}, {}}, nil},
		{"no_nodes", nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			streams := make([]<-chan int, 0, len(tc.nodes))
			for _, samples := range tc.nodes {
				streams = append(streams, send(samples...))
			}
			got := MergeMetrics(streams...)
			if len(got) != len(tc.want) {
				t.Fatalf("MergeMetrics(%v) = %v, want %v", tc.nodes, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("MergeMetrics(%v) = %v, want %v", tc.nodes, got, tc.want)
				}
			}
		})
	}
}
