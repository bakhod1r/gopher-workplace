package modulo

import (
	"errors"
	"testing"
)

func TestMod(t *testing.T) {
	cases := []struct {
		name    string
		a, b    int
		want    int
		wantErr error
	}{
		{"basic", 10, 3, 1, nil},
		{"exact", 9, 3, 0, nil},
		{"negative_dividend", -7, 3, -1, nil},
		{"zero_modulus", 5, 0, 0, ErrZeroModulus},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Mod(tc.a, tc.b)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Mod(%d, %d) err = %v, want %v", tc.a, tc.b, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("Mod(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
