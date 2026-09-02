package equalslices

import "testing"

func TestEqual(t *testing.T) {
	cases := []struct {
		name string
		a, b []int
		want bool
	}{
		{"same", []int{1, 2}, []int{1, 2}, true},
		{"reordered", []int{1, 2}, []int{2, 1}, false},
		{"different_length", []int{1}, []int{1, 2}, false},
		{"both_empty", []int{}, []int{}, true},
		{"nil_and_empty", nil, []int{}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Equal(tc.a, tc.b); got != tc.want {
				t.Errorf("Equal(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
	if !Equal([]string{"a"}, []string{"a"}) {
		t.Error(`Equal([]string{"a"}, []string{"a"}) = false, want true`)
	}
}
