package lookupkey

import (
	"errors"
	"testing"
)

func TestLookup(t *testing.T) {
	m := map[string]int{"a": 1, "zero": 0}

	cases := []struct {
		name    string
		m       map[string]int
		key     string
		want    int
		wantErr error
	}{
		{"present", m, "a", 1, nil},
		{"stored_zero", m, "zero", 0, nil},
		{"absent", m, "z", 0, ErrNotFound},
		{"nil_map", nil, "a", 0, ErrNotFound},
		{"empty_map", map[string]int{}, "a", 0, ErrNotFound},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Lookup(tc.m, tc.key)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Lookup(%v, %q) err = %v, want %v", tc.m, tc.key, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("Lookup(%v, %q) = %d, want %d", tc.m, tc.key, got, tc.want)
			}
		})
	}
}
