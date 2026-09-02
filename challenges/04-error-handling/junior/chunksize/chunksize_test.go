package chunksize

import (
	"errors"
	"testing"
)

func TestChunkSize(t *testing.T) {
	cases := []struct {
		name         string
		total, parts int
		want         int
		wantErr      error
	}{
		{"rounds_up", 10, 3, 4, nil},
		{"exact", 9, 3, 3, nil},
		{"single_part", 7, 1, 7, nil},
		{"zero_total", 0, 3, 0, nil},
		{"zero_parts", 10, 0, 0, ErrBadParts},
		{"negative_parts", 10, -2, 0, ErrBadParts},
		{"negative_total", -1, 3, 0, ErrNegativeTotal},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ChunkSize(tc.total, tc.parts)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("ChunkSize(%d, %d) err = %v, want %v", tc.total, tc.parts, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("ChunkSize(%d, %d) = %d, want %d", tc.total, tc.parts, got, tc.want)
			}
		})
	}
}
