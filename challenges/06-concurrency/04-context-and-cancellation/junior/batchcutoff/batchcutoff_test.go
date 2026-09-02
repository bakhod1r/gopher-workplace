package batchcutoff

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestMissedCutoff(t *testing.T) {
	cases := []struct {
		name   string
		cutoff time.Time
	}{
		{"one_hour_ago", time.Now().Add(-time.Hour)},
		{"one_nanosecond_ago", time.Now().Add(-time.Nanosecond)},
		{"unix_epoch", time.Unix(0, 0)},
		{"year_2000", time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{"zero_time", time.Time{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := MissedCutoff(tc.cutoff)
			if got != context.DeadlineExceeded {
				t.Errorf("MissedCutoff(%v) = %v, want %v", tc.cutoff, got, context.DeadlineExceeded)
			}
			if errors.Is(got, context.Canceled) {
				t.Error("got Canceled, want DeadlineExceeded")
			}
		})
	}
}
