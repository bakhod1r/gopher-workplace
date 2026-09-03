package failurecause

import (
	"context"
	"errors"
	"testing"
)

func rows(n int) []string {
	out := make([]string, n)
	for i := range out {
		out[i] = "row"
	}
	return out
}

func TestExport(t *testing.T) {
	cases := []struct {
		name      string
		rows      int
		quota     int
		wantCount int
		wantErr   error
	}{
		{"under_quota", 5, 10, 5, nil},
		{"exactly_quota", 3, 3, 3, nil},
		{"over_quota", 5, 3, 3, ErrQuotaExceeded},
		{"no_rows", 0, 3, 0, nil},
		{"zero_quota", 2, 0, 0, ErrQuotaExceeded},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Export(context.Background(), rows(tc.rows), tc.quota)
			if got != tc.wantCount || !errors.Is(err, tc.wantErr) {
				t.Errorf("Export() = %d, %v; want %d, %v", got, err, tc.wantCount, tc.wantErr)
			}
		})
	}
}

func TestCancelledParentIsReportedAsCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	got, err := Export(ctx, rows(5), 10)
	if got != 0 || !errors.Is(err, context.Canceled) {
		t.Errorf("Export() = %d, %v; want 0, context.Canceled", got, err)
	}
}

func TestQuotaCauseIsNotBareCanceled(t *testing.T) {
	_, err := Export(context.Background(), rows(5), 2)
	if errors.Is(err, context.Canceled) {
		t.Error("quota failure was reported as context.Canceled; the cause is lost")
	}
	if !errors.Is(err, ErrQuotaExceeded) {
		t.Errorf("err = %v, want ErrQuotaExceeded", err)
	}
}

func TestExpiredDeadlineIsReported(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()

	got, err := Export(ctx, rows(3), 10)
	if got != 0 || !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("Export() = %d, %v; want 0, context.DeadlineExceeded", got, err)
	}
}
