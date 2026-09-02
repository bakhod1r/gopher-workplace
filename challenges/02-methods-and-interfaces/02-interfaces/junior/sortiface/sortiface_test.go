package sortiface

import "testing"

func TestIntSliceMethods(t *testing.T) {
	s := IntSlice{3, 1}
	if s.Len() != 2 {
		t.Errorf("Len = %d, want 2", s.Len())
	}
	if !s.Less(1, 0) {
		t.Error("Less(1, 0) = false, want true")
	}
	s.Swap(0, 1)
	if s[0] != 1 || s[1] != 3 {
		t.Errorf("after Swap = %v, want [1 3]", s)
	}
}

func TestBubbleSort(t *testing.T) {
	cases := []struct {
		name string
		in   IntSlice
		want []int
	}{
		{"unsorted", IntSlice{3, 1, 2}, []int{1, 2, 3}},
		{"sorted", IntSlice{1, 2, 3}, []int{1, 2, 3}},
		{"reverse", IntSlice{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"dupes", IntSlice{2, 1, 2}, []int{1, 2, 2}},
		{"empty", IntSlice{}, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			BubbleSort(tc.in)
			for i := range tc.want {
				if tc.in[i] != tc.want[i] {
					t.Fatalf("got %v, want %v", tc.in, tc.want)
				}
			}
		})
	}
}
