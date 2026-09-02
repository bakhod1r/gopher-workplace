package exportrows

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func cancelled() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func expired() context.Context {
	ctx, cancel := context.WithDeadline(context.Background(), time.Unix(0, 0))
	cancel()
	return ctx
}

func TestExportRows(t *testing.T) {
	live, cancelLive := context.WithCancel(context.Background())
	defer cancelLive()

	cases := []struct {
		name    string
		ctx     context.Context
		ids     []int
		want    []string
		wantErr error
	}{
		{"two_rows", live, []int{1, 2}, []string{"row-1", "row-2"}, nil},
		{"single_row", live, []int{42}, []string{"row-42"}, nil},
		{"no_rows", live, nil, []string{}, nil},
		{"empty_slice", live, []int{}, []string{}, nil},
		{"many_rows", live, []int{1, 2, 3, 4, 5}, []string{"row-1", "row-2", "row-3", "row-4", "row-5"}, nil},
		{"user_cancelled", cancelled(), []int{1, 2}, nil, context.Canceled},
		{"budget_expired", expired(), []int{1, 2}, nil, context.DeadlineExceeded},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ExportRows(tc.ctx, tc.ids)
			if err != tc.wantErr {
				t.Fatalf("ExportRows() err = %v, want %v", err, tc.wantErr)
			}
			if err != nil {
				if got != nil {
					t.Errorf("ExportRows() = %v, want nil on error", got)
				}
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ExportRows() = %v, want %v", got, tc.want)
			}
		})
	}
}
