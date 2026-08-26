package isempty

import "testing"

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		name string
		s    Stack
		want bool
	}{
		{"nil_items", Stack{}, true},
		{"empty_slice", Stack{items: []int{}}, true},
		{"one_item", Stack{items: []int{1}}, false},
		{"many_items", Stack{items: []int{1, 2, 3}}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.s.IsEmpty(); got != tc.want {
				t.Errorf("Stack{len=%d}.IsEmpty() = %v, want %v",
					len(tc.s.items), got, tc.want)
			}
		})
	}
}
