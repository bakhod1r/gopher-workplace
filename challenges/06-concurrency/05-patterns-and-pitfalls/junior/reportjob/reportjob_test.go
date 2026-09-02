package reportjob

import (
	"context"
	"errors"
	"testing"
)

func TestRunReport(t *testing.T) {
	cases := []struct {
		name    string
		rows    []int
		cancel  bool
		want    int
		wantErr error
	}{
		{"three_rows", []int{1, 2, 3}, false, 6, nil},
		{"single_row", []int{42}, false, 42, nil},
		{"negative_rows", []int{-5, 5, -1}, false, -1, nil},
		{"no_rows", nil, false, 0, nil},
		{"cancelled_before_start", []int{1, 2, 3}, true, 0, context.Canceled},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if tc.cancel {
				cancel()
			}

			got, err := RunReport(ctx, tc.rows)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("RunReport(%v) err = %v, want %v", tc.rows, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("RunReport(%v) = %d, want %d", tc.rows, got, tc.want)
			}
		})
	}
}
