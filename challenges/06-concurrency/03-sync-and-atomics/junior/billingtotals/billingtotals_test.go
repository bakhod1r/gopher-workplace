package billingtotals

import (
	"sync"
	"testing"
)

func TestTotals(t *testing.T) {
	cases := []struct {
		name    string
		amounts []int64
		wantSum int64
		wantN   int64
		wantAvg int64
	}{
		{"empty_run", nil, 0, 0, 0},
		{"single_invoice", []int64{300}, 300, 1, 300},
		{"two_invoices", []int64{300, 100}, 400, 2, 200},
		{"truncated_mean", []int64{10, 11}, 21, 2, 10},
		{"with_credit_note", []int64{100, -40}, 60, 2, 30},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var tt Totals
			for _, a := range tc.amounts {
				tt.Add(a)
			}
			if got := tt.Total(); got != tc.wantSum {
				t.Errorf("Total() = %d, want %d", got, tc.wantSum)
			}
			if got := tt.Count(); got != tc.wantN {
				t.Errorf("Count() = %d, want %d", got, tc.wantN)
			}
			if got := tt.Average(); got != tc.wantAvg {
				t.Errorf("Average() = %d, want %d", got, tc.wantAvg)
			}
		})
	}
}

func TestTotalsConcurrent(t *testing.T) {
	var tt Totals
	const workers = 8
	const per = 250
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < per; j++ {
				tt.Add(50)
				tt.Average()
			}
		}()
	}
	wg.Wait()

	if got, want := tt.Count(), int64(workers*per); got != want {
		t.Errorf("Count() = %d, want %d", got, want)
	}
	if got, want := tt.Total(), int64(workers*per*50); got != want {
		t.Errorf("Total() = %d, want %d", got, want)
	}
	if got := tt.Average(); got != 50 {
		t.Errorf("Average() = %d, want 50", got)
	}
}
