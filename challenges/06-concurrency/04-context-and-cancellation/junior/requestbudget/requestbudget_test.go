package requestbudget

import (
	"context"
	"testing"
	"time"
)

func TestBudget(t *testing.T) {
	at := time.Date(2030, time.March, 1, 12, 0, 0, 0, time.UTC)

	withDeadline, cancelDeadline := context.WithDeadline(context.Background(), at)
	defer cancelDeadline()

	withCancel, cancelCancel := context.WithCancel(context.Background())
	defer cancelCancel()

	child, cancelChild := context.WithCancel(withDeadline)
	defer cancelChild()

	cases := []struct {
		name   string
		ctx    context.Context
		want   time.Time
		wantOK bool
	}{
		{"process_root", context.Background(), time.Time{}, false},
		{"todo", context.TODO(), time.Time{}, false},
		{"only_a_trace_value", context.WithValue(context.Background(), "trace", "abc"), time.Time{}, false},
		{"only_cancellable", withCancel, time.Time{}, false},
		{"has_deadline", withDeadline, at, true},
		{"child_inherits_budget", child, at, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := Budget(tc.ctx)
			if ok != tc.wantOK {
				t.Fatalf("Budget() ok = %v, want %v", ok, tc.wantOK)
			}
			if !got.Equal(tc.want) {
				t.Errorf("Budget() time = %v, want %v", got, tc.want)
			}
		})
	}
}
