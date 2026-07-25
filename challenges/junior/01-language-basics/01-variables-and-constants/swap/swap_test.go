package swap

import "testing"

func TestSwap(t *testing.T) {
	cases := []struct {
		name  string
		a, b  int
		wantA int
		wantB int
	}{
		{"basic", 1, 2, 2, 1},
		{"equal", 9, 9, 9, 9},
		{"negative", -3, 7, 7, -3},
		{"zero", 0, 5, 5, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotA, gotB := Swap(tc.a, tc.b)
			if gotA != tc.wantA || gotB != tc.wantB {
				t.Errorf("Swap(%d, %d) = %d, %d; want %d, %d",
					tc.a, tc.b, gotA, gotB, tc.wantA, tc.wantB)
			}
		})
	}
}
