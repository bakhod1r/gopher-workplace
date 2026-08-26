package reduceslice

import "testing"

func TestProduct(t *testing.T) {
	cases := []struct {
		name string
		list IntList
		want int
	}{
		{"mixed", IntList{2, 3, 4}, 24},
		{"with_zero", IntList{5, 0, 10}, 0},
		{"negative", IntList{-2, 3}, -6},
		{"empty", IntList{}, 1},
		{"nil", nil, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.list.Product(); got != tc.want {
				t.Errorf("Product() = %d, want %d", got, tc.want)
			}
		})
	}
}
