package mergesort

import "testing"

func feed(vs ...int) *SliceFeed { return &SliceFeed{Data: vs} }

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

func TestSliceFeed(t *testing.T) {
	f := feed(1, 2)
	if v, ok := f.Peek(); v != 1 || !ok {
		t.Errorf("Peek = %d, %v", v, ok)
	}
	if v, ok := f.Peek(); v != 1 || !ok {
		t.Error("Peek must not consume")
	}
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

func TestMerge(t *testing.T) {
	cases := []struct {
		name string
		a, b []int
		want []int
	}{
		{"interleaved", []int{1, 3}, []int{2, 4}, []int{1, 2, 3, 4}},
		{"a_empty", nil, []int{1}, []int{1}},
		{"b_empty", []int{1}, nil, []int{1}},
		{"both_empty", nil, nil, nil},
		{"duplicates", []int{1, 1}, []int{1}, []int{1, 1, 1}},
		{"disjoint", []int{5, 6}, []int{1, 2}, []int{1, 2, 5, 6}},
		{"uneven", []int{1}, []int{2, 3, 4}, []int{1, 2, 3, 4}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Merge(feed(tc.a...), feed(tc.b...))
			if !eq(got, tc.want) {
				t.Errorf("Merge = %v, want %v", got, tc.want)
			}
		})
	}
}
