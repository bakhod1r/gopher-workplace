package revenuereport

import (
	"context"
	"testing"
)

func closedRows(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func cancelled() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func expired() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	cancel()
	return ctx
}

func TestTotalRevenue(t *testing.T) {
	live, cancelLive := context.WithCancel(context.Background())
	defer cancelLive()

	cases := []struct {
		name    string
		ctx     context.Context
		rows    <-chan int
		want    int
		wantErr error
	}{
		{"three_rows", live, closedRows(100, 250, 25), 375, nil},
		{"single_row", live, closedRows(4200), 4200, nil},
		{"no_rows", live, closedRows(), 0, nil},
		{"refunds_included", live, closedRows(500, -200), 300, nil},
		{"user_cancelled_export", cancelled(), make(chan int), 0, context.Canceled},
		{"export_budget_expired", expired(), make(chan int), 0, context.DeadlineExceeded},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := TotalRevenue(tc.ctx, tc.rows)
			if err != tc.wantErr {
				t.Fatalf("TotalRevenue() err = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("TotalRevenue() = %d, want %d", got, tc.want)
			}
		})
	}
}
