package typeassert

import "testing"

func TestAsInt(t *testing.T) {
	cases := []struct {
		name string
		in   any
		want int
		ok   bool
	}{
		{"int", 5, 5, true},
		{"zero", 0, 0, true},
		{"string", "5", 0, false},
		{"float", 5.0, 0, false},
		{"nil", nil, 0, false},
		{"int64", int64(5), 0, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := AsInt(tc.in)
			if got != tc.want || ok != tc.ok {
				t.Errorf("AsInt(%#v) = %d, %v; want %d, %v", tc.in, got, ok, tc.want, tc.ok)
			}
		})
	}
}

func TestSumInts(t *testing.T) {
	cases := []struct {
		name string
		vs   []any
		want int
	}{
		{"mixed", []any{1, "x", 2}, 3},
		{"all_ints", []any{1, 2, 3}, 6},
		{"none", []any{"a", 1.5, nil}, 0},
		{"empty", nil, 0},
		{"negatives", []any{-1, 1}, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SumInts(tc.vs); got != tc.want {
				t.Errorf("SumInts = %d, want %d", got, tc.want)
			}
		})
	}
}
