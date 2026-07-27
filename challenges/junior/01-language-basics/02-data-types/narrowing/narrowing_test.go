package narrowing

import "testing"

func TestToInt32(t *testing.T) {
	cases := []struct {
		name string
		in   int64
		want int32
	}{
		{"small", 42, 42},
		{"negative", -7, -7},
		{"int32 max", 2147483647, 2147483647},
		{"one past max wraps", 2147483648, -2147483648},
		{"2^32 wraps to zero", 4294967296, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToInt32(tc.in); got != tc.want {
				t.Errorf("ToInt32(%d) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}
