package positive

import (
	"errors"
	"testing"
)

func TestPositive(t *testing.T) {
	cases := []struct {
		name    string
		n       int
		want    int
		wantErr error
	}{
		{"positive", 5, 5, nil},
		{"one", 1, 1, nil},
		{"zero", 0, 0, ErrNotPositive},
		{"negative", -3, 0, ErrNotPositive},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Positive(tc.n)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Positive(%d) err = %v, want %v", tc.n, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("Positive(%d) = %d, want %d", tc.n, got, tc.want)
			}
		})
	}
}
