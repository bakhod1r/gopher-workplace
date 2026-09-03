package ingestpipeline

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"
)

var errBadRecord = errors.New("record rejected by schema")

// upper normalises a record, rejecting anything named "bad". It also fails the
// test if it is ever handed a context that is already finished.
func upper(t *testing.T, calls *int) Normalise {
	return func(ctx context.Context, record string) (string, error) {
		t.Helper()
		*calls++
		if err := ctx.Err(); err != nil {
			t.Errorf("normalise called with a finished context: %v", err)
		}
		if record == "bad" {
			return "", errBadRecord
		}
		return strings.ToUpper(record), nil
	}
}

func uploadAborted() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func budgetExpired() context.Context {
	ctx, cancel := context.WithDeadline(context.Background(), time.Unix(0, 0))
	cancel()
	return ctx
}

func TestRunIngestPipeline(t *testing.T) {
	cases := []struct {
		name      string
		ctx       context.Context
		records   []string
		wantClean string
		wantErr   error
		wantCalls int
	}{
		{"no_records", context.Background(), nil, "", nil, 0},
		{"single_record", context.Background(), []string{"alpha"}, "ALPHA", nil, 1},
		{"three_records_in_order", context.Background(), []string{"alpha", "beta", "gamma"}, "ALPHA,BETA,GAMMA", nil, 3},
		{"second_record_rejected", context.Background(), []string{"alpha", "bad", "gamma"}, "", errBadRecord, 2},
		{"first_record_rejected", context.Background(), []string{"bad", "alpha"}, "", errBadRecord, 1},
		{"upload_aborted", uploadAborted(), []string{"alpha", "beta"}, "", context.Canceled, 0},
		{"ingest_budget_expired", budgetExpired(), []string{"alpha"}, "", context.DeadlineExceeded, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			clean, err := RunIngestPipeline(tc.ctx, tc.records, upper(t, &calls))
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("RunIngestPipeline() error = %v, want %v", err, tc.wantErr)
			}
			if got := strings.Join(clean, ","); got != tc.wantClean {
				t.Errorf("clean = %q, want %q", got, tc.wantClean)
			}
			if calls != tc.wantCalls {
				t.Errorf("normalise called %d times, want %d", calls, tc.wantCalls)
			}
		})
	}
}
