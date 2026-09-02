package cfgget

import (
	"errors"
	"testing"
)

func TestGet(t *testing.T) {
	cfg := map[string]string{"port": "80", "blank": ""}

	cases := []struct {
		name    string
		cfg     map[string]string
		key     string
		want    string
		wantErr error
	}{
		{"present", cfg, "port", "80", nil},
		{"stored_empty", cfg, "blank", "", nil},
		{"absent", cfg, "host", "", ErrMissingKey},
		{"nil_map", nil, "port", "", ErrMissingKey},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Get(tc.cfg, tc.key)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Get(%v, %q) err = %v, want %v", tc.cfg, tc.key, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("Get(%v, %q) = %q, want %q", tc.cfg, tc.key, got, tc.want)
			}
		})
	}
}
