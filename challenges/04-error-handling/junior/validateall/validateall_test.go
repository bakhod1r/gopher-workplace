package validateall

import "testing"

func TestValidate(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"all_valid", []int{1, 2, 0}, 0},
		{"two_bad", []int{1, -2, -3}, 2},
		{"all_bad", []int{-1, -2}, 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Validate(tc.nums)
			if len(got) != tc.want {
				t.Fatalf("Validate(%v) returned %d errors, want %d", tc.nums, len(got), tc.want)
			}
			if tc.want == 0 && got != nil {
				t.Errorf("Validate(%v) = %v, want nil", tc.nums, got)
			}
			for i, err := range got {
				if err != ErrNegative {
					t.Errorf("Validate(%v)[%d] = %v, want ErrNegative", tc.nums, i, err)
				}
			}
		})
	}
}
