package mergemany

import "testing"

func feed(vs ...int) *SortedFeed { return &SortedFeed{Data: vs} }

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

func TestSortedFeed(t *testing.T) {
	f := feed(1, 2)
	if v, ok := f.Next(); v != 1 || !ok {
		t.Errorf("Next = %d, %v", v, ok)
	}
	if v, ok := f.Next(); v != 2 || !ok {
		t.Errorf("Next = %d, %v", v, ok)
	}
	if _, ok := f.Next(); ok {
		t.Error("drained Next should report false")
	}
}

func TestMergeAll(t *testing.T) {
	cases := []struct {
		name  string
		feeds []Feed
		want  []int
	}{
		{"three", []Feed{feed(1, 4), feed(2, 5), feed(3)}, []int{1, 2, 3, 4, 5}},
		{"one", []Feed{feed(1, 2)}, []int{1, 2}},
		{"with_empty", []Feed{feed(), feed(1), feed()}, []int{1}},
		{"all_empty", []Feed{feed(), feed()}, nil},
		{"none", nil, nil},
		{"duplicates", []Feed{feed(1, 1), feed(1)}, []int{1, 1, 1}},
		{"disjoint", []Feed{feed(5, 6), feed(1, 2)}, []int{1, 2, 5, 6}},
		{"negatives", []Feed{feed(-5, 0), feed(-3)}, []int{-5, -3, 0}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := MergeAll(tc.feeds...); !eq(got, tc.want) {
				t.Errorf("MergeAll = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestManyFeeds(t *testing.T) {
	const k = 50
	feeds := make([]Feed, k)
	for i := range feeds {
		data := make([]int, 0, 20)
		for j := 0; j < 20; j++ {
			data = append(data, i+j*k)
		}
		feeds[i] = &SortedFeed{Data: data}
	}

	got := MergeAll(feeds...)
	if len(got) != k*20 {
		t.Fatalf("len = %d, want %d", len(got), k*20)
	}
	for i := 1; i < len(got); i++ {
		if got[i] < got[i-1] {
			t.Fatalf("output not sorted at %d: %d < %d", i, got[i], got[i-1])
		}
	}
}

func BenchmarkMergeAll(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		feeds := make([]Feed, 8)
		for j := range feeds {
			data := make([]int, 0, 100)
			for k := 0; k < 100; k++ {
				data = append(data, j+k*8)
			}
			feeds[j] = &SortedFeed{Data: data}
		}
		MergeAll(feeds...)
	}
}
