package topkstream

import "testing"

func eq(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range want {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func TestTopK(t *testing.T) {
	cases := []struct {
		name string
		k    int
		data []int
		want []int
	}{
		{"basic", 2, []int{1, 5, 3}, []int{5, 3}},
		{"k_exceeds_input", 3, []int{1, 2}, []int{2, 1}},
		{"k_zero", 0, []int{1, 2}, nil},
		{"descending_input", 2, []int{9, 8, 7}, []int{9, 8}},
		{"ascending_input", 2, []int{1, 2, 3}, []int{3, 2}},
		{"duplicates", 3, []int{5, 5, 5, 1}, []int{5, 5, 5}},
		{"negatives", 2, []int{-5, -1, -3}, []int{-1, -3}},
		{"empty", 2, nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Stream(&SliceSource{Data: tc.data}, NewTopK(tc.k))
			if !eq(got, tc.want) {
				t.Errorf("Stream = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMemoryBoundedByK(t *testing.T) {
	const n = 200000
	data := make([]int, 0, n)
	for i := 0; i < n; i++ {
		data = append(data, i%9973)
	}

	agg := NewTopK(10)
	got := Stream(&SliceSource{Data: data}, agg)

	if len(agg.vals) > 10 {
		t.Errorf("window holds %d values, want at most 10", len(agg.vals))
	}
	if cap(agg.vals) > 10 {
		t.Errorf("window capacity grew to %d, want at most 10", cap(agg.vals))
	}
	if len(got) != 10 || got[0] != 9972 {
		t.Errorf("Result = %v", got)
	}
}

func TestDescendingOrder(t *testing.T) {
	got := Stream(&SliceSource{Data: []int{3, 1, 4, 1, 5, 9, 2, 6}}, NewTopK(4))
	if !eq(got, []int{9, 6, 5, 4}) {
		t.Errorf("Stream = %v, want [9 6 5 4]", got)
	}
}

func BenchmarkStream(b *testing.B) {
	data := make([]int, 10000)
	for i := range data {
		data[i] = i % 977
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Stream(&SliceSource{Data: data}, NewTopK(10))
	}
}
